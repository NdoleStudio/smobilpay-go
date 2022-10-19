package smobilpay

import (
	"net/http"
	"time"
)

type clientConfig struct {
	httpClient                  *http.Client
	collectSyncVerifyInterval   time.Duration
	collectSyncVerifyRetryCount uint
	baseURL                     string
	accessToken                 string
	accessSecret                string
}

func defaultClientConfig() *clientConfig {
	return &clientConfig{
		httpClient:                  http.DefaultClient,
		collectSyncVerifyInterval:   20 * time.Second,
		collectSyncVerifyRetryCount: 15,
		baseURL:                     "https://s3p.smobilpay.staging.maviance.info/v2",
	}
}
