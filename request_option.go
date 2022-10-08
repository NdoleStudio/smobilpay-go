package smobilpay

import (
	"time"
)

// RequestOption is options for constructing an API request
type RequestOption interface {
	apply(config *requestConfig)
}

type requestOptionFunc func(config *requestConfig)

func (fn requestOptionFunc) apply(config *requestConfig) {
	fn(config)
}

// WithTimestamp sets the timestamp in the s3pAuth_timestamp header
func WithTimestamp(timestamp time.Time) RequestOption {
	return requestOptionFunc(func(config *requestConfig) {
		config.timestamp = timestamp
	})
}

// WithNonce sets the nonce in the s3pAuth_nonce header
func WithNonce(nonce string) RequestOption {
	return requestOptionFunc(func(config *requestConfig) {
		config.nonce = nonce
	})
}
