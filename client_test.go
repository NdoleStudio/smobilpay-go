package smobilpay

import (
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/NdoleStudio/smobilpay-go/internal/helpers"
	"github.com/NdoleStudio/smobilpay-go/internal/stubs"
	"github.com/stretchr/testify/assert"
)

func TestClient_Ping_Ok(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.PingOk())
	accessToken := "6B352110-4716-11ED-963F-0800200C9A66"
	client := New(
		WithBaseURL(server.URL),
		WithAccessToken(accessToken),
		WithAccessSecret("1B875FB0-4717-11ED-963F-0800200C9A66"),
	)
	nonce := "95cdf110-4614-4d95-b6c2-f14fe01c4995"

	// Act
	status, response, err := client.Ping(
		context.Background(),
		WithRequestTimestamp(time.Now()),
		WithRequestNonce(nonce),
	)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, nonce, status.Nonce)
	assert.Equal(t, accessToken, status.Key)

	// Teardown
	server.Close()
}

func TestClient_Ping_Error(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusBadRequest, stubs.PingError())
	accessToken := "6B352110-4716-11ED-963F-0800200C9A66"
	client := New(
		WithBaseURL(server.URL),
		WithAccessToken(accessToken),
		WithAccessSecret("1B875FB0-4717-11ED-963F-0800200C9A66"),
	)
	nonce := "95cdf110-4614-4d95-b6c2-f14fe01c4995"

	// Act
	status, response, err := client.Ping(
		context.Background(),
		WithRequestTimestamp(time.Now()),
		WithRequestNonce(nonce),
	)

	// Assert
	assert.NotNil(t, err)
	assert.Nil(t, status)
	assert.Equal(t, http.StatusBadRequest, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}

func TestClient_Ping_Request(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	request := new(http.Request)
	server := helpers.MakeRequestCapturingTestServer(http.StatusOK, stubs.PingOk(), request)
	accessToken := "6B352110-4716-11ED-963F-0800200C9A66"
	accessSecret := "1B875FB0-4717-11ED-963F-0800200C9A66"
	client := New(
		WithBaseURL(server.URL),
		WithAccessToken(accessToken),
		WithAccessSecret(accessSecret),
	)
	nonce := "95cdf110-4614-4d95-b6c2-f14fe01c4995"
	signature := computeHmac(
		"GET&"+url.QueryEscape(server.URL)+"%2Fping&s3pAuth_nonce%3D95cdf110-4614-4d95-b6c2-f14fe01c4995%26s3pAuth_signature_method%3DHMAC-SHA1%26s3pAuth_timestamp%3D1613869830%26s3pAuth_token%3D6B352110-4716-11ED-963F-0800200C9A66",
		accessSecret,
	)

	// Act
	_, _, err := client.Ping(
		context.Background(),
		WithRequestTimestamp(time.Date(2021, time.Month(2), 21, 1, 10, 30, 0, time.UTC)),
		WithRequestNonce(nonce),
	)

	// Assert
	assert.Nil(t, err)
	assert.Equal(
		t,
		"s3pAuth,s3pAuth_nonce=\"95cdf110-4614-4d95-b6c2-f14fe01c4995\",s3pAuth_signature=\""+signature+"\",s3pAuth_signature_method=\"HMAC-SHA1\",s3pAuth_timestamp=\"1613869830\",s3pAuth_token=\"6B352110-4716-11ED-963F-0800200C9A66\"",
		request.Header.Get("Authorization"),
	)

	// Teardown
	server.Close()
}

func computeHmac(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha1.New, key)
	_, _ = h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
