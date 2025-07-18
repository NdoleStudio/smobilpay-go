package helpers

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
)

// MakeTestServer creates an api server for testing
func MakeTestServer(responseCode int, body []byte) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(responseCode)
		_, err := res.Write(body)
		if err != nil {
			panic(err)
		}
	}))
}

// MakeTestServerWithMultipleResponses creates an api server for testing with multiple responses
func MakeTestServerWithMultipleResponses(responseCode int, responses [][]byte) *httptest.Server {
	count := 0
	return httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(responseCode)
		_, err := res.Write(responses[count])
		if err != nil {
			panic(err)
		}
		count++
	}))
}

// MakeRequestCapturingTestServer creates an api server that captures the request object
func MakeRequestCapturingTestServer(responseCode int, response []byte, request *http.Request) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(responseWriter http.ResponseWriter, req *http.Request) {
		clonedRequest := req.Clone(context.Background())

		// clone body
		body, err := io.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}
		req.Body = io.NopCloser(bytes.NewReader(body))
		clonedRequest.Body = io.NopCloser(bytes.NewReader(body))

		*request = *clonedRequest

		responseWriter.WriteHeader(responseCode)
		_, err = responseWriter.Write(response)
		if err != nil {
			panic(err)
		}
	}))
}
