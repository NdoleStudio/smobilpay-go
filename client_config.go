package client

import "net/http"

type clientConfig struct {
	httpClient   *http.Client
	baseURL      string
	accessToken  string
	accessSecret string
}

func defaultClientConfig() *clientConfig {
	return &clientConfig{
		httpClient: http.DefaultClient,
		baseURL:    "https://s3p.smobilpay.staging.maviance.info/v2",
	}
}
