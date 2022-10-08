package smobilpay

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type service struct {
	client *Client
}

// Client is the campay API client.
// Do not instantiate this client with Client{}. Use the New method instead.
type Client struct {
	httpClient   *http.Client
	common       service
	baseURL      string
	accessToken  string
	accessSecret string

	Status *topupService
}

// New creates and returns a new campay.Client from a slice of campay.ClientOption.
func New(options ...Option) *Client {
	config := defaultClientConfig()

	for _, option := range options {
		option.apply(config)
	}

	client := &Client{
		httpClient:   config.httpClient,
		accessToken:  config.accessToken,
		accessSecret: config.accessSecret,
		baseURL:      config.baseURL,
	}

	client.common.client = client
	client.Status = (*topupService)(&client.common)
	return client
}

// Ping checks if the API is available
//
// https://apidocs.smobilpay.com/s3papi/API-Reference.2066448558.html
func (client *Client) Ping(ctx context.Context, options ...RequestOption) (*PingStatus, *Response, error) {
	request, err := client.newRequest(ctx, options, http.MethodGet, "/ping", nil)
	if err != nil {
		return nil, nil, err
	}

	response, err := client.do(request)
	if err != nil {
		return nil, response, err
	}

	status := new(PingStatus)
	if err = json.Unmarshal(*response.Body, status); err != nil {
		return nil, response, err
	}

	return status, response, nil
}

func (client *Client) makeRequestConfig(options []RequestOption) *requestConfig {
	config := defaultRequestConfig()

	for _, option := range options {
		option.apply(config)
	}

	return config
}

// newRequest creates an API request. A relative URL can be provided in uri,
// in which case it is resolved relative to the BaseURL of the Client.
// URI's should always be specified without a preceding slash.
func (client *Client) newRequest(ctx context.Context, options []RequestOption, method, uri string, body map[string]string) (*http.Request, error) {
	config := client.makeRequestConfig(options)

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, client.baseURL+uri, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", client.getAuthHeader(req, config, client.authPayload(req, body)))

	return req, nil
}

func (client *Client) authPayload(request *http.Request, body map[string]string) map[string]string {
	if request.Method != http.MethodGet {
		return body
	}

	payload := map[string]string{}
	for key, value := range request.URL.Query() {
		payload[key] = value[0]
	}

	return payload
}

func (client *Client) getBaseHmacAuthString(request *http.Request) string {
	return fmt.Sprintf(
		"%s&%s",
		strings.ToUpper(request.Method),
		url.QueryEscape(
			fmt.Sprintf("%s://%s%s", request.URL.Scheme, request.URL.Host, request.URL.Path),
		),
	)
}

func (client *Client) getPayloadHmacAuthString(config *requestConfig, payload map[string]string) string {
	params := url.Values{}
	params.Add("s3pAuth_nonce", config.nonce)
	params.Add("s3pAuth_signature_method", "HMAC-SHA1")
	params.Add("s3pAuth_timestamp", config.timestampString())
	params.Add("s3pAuth_token", client.accessToken)
	for key, value := range payload {
		params.Add(key, value)
	}

	return params.Encode()
}

func (client *Client) getAuthHeader(request *http.Request, config *requestConfig, payload map[string]string) string {
	return fmt.Sprintf(
		"s3pAuth,s3pAuth_nonce=\"%s\",s3pAuth_signature=\"%s\",s3pAuth_signature_method=\"HMAC-SHA1\",s3pAuth_timestamp=\"%s\",s3pAuth_token=\"%s\"",
		config.nonce,
		client.computeHmac(client.getBaseHmacAuthString(request)+"&"+url.QueryEscape(client.getPayloadHmacAuthString(config, payload))),
		config.timestampString(),
		client.accessToken,
	)
}

func (client *Client) computeHmac(message string) string {
	log.Println(message)
	key := []byte(client.accessSecret)
	h := hmac.New(sha1.New, key)
	_, _ = h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// do carries out an HTTP request and returns a Response
func (client *Client) do(req *http.Request) (*Response, error) {
	if req == nil {
		return nil, fmt.Errorf("%T cannot be nil", req)
	}

	httpResponse, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() { _ = httpResponse.Body.Close() }()

	resp, err := client.newResponse(httpResponse)
	if err != nil {
		return resp, err
	}

	_, err = io.Copy(ioutil.Discard, httpResponse.Body)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// newResponse converts an *http.Response to *Response
func (client *Client) newResponse(httpResponse *http.Response) (*Response, error) {
	if httpResponse == nil {
		return nil, fmt.Errorf("%T cannot be nil", httpResponse)
	}

	resp := new(Response)
	resp.HTTPResponse = httpResponse

	buf, err := ioutil.ReadAll(resp.HTTPResponse.Body)
	if err != nil {
		return nil, err
	}
	resp.Body = &buf

	return resp, resp.Error()
}
