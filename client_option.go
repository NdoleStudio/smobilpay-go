package smobilpay

import (
	"net/http"
	"strings"
	"time"
)

// Option is options for constructing a client
type Option interface {
	apply(config *clientConfig)
}

type clientOptionFunc func(config *clientConfig)

func (fn clientOptionFunc) apply(config *clientConfig) {
	fn(config)
}

// WithHTTPClient sets the underlying HTTP client used for API requests.
// By default, http.DefaultClient is used.
func WithHTTPClient(httpClient *http.Client) Option {
	return clientOptionFunc(func(config *clientConfig) {
		if httpClient != nil {
			config.httpClient = httpClient
		}
	})
}

// WithBaseURL set's the base url for the smobilpay API
func WithBaseURL(baseURL string) Option {
	return clientOptionFunc(func(config *clientConfig) {
		if baseURL != "" {
			config.baseURL = strings.TrimRight(baseURL, "/")
		}
	})
}

// WithAccessToken sets the access token for the smobilpay api
func WithAccessToken(accessToken string) Option {
	return clientOptionFunc(func(config *clientConfig) {
		config.accessToken = accessToken
	})
}

// WithAccessSecret sets the access secret for the smobilpay api
func WithAccessSecret(accessSecret string) Option {
	return clientOptionFunc(func(config *clientConfig) {
		config.accessSecret = accessSecret
	})
}

// WithCollectSyncVerifyInterval sets the interval for calling the `/verifytx` endpoint to check the status of pending transactions
func WithCollectSyncVerifyInterval(interval time.Duration) Option {
	return clientOptionFunc(func(config *clientConfig) {
		config.collectSyncVerifyInterval = interval
	})
}

// WithCollectSyncVerifyRetryCount sets the number of retries for calling the `/verifytx` endpoint to check the status of pending transactions
func WithCollectSyncVerifyRetryCount(retryCount uint) Option {
	return clientOptionFunc(func(config *clientConfig) {
		config.collectSyncVerifyRetryCount = retryCount
	})
}
