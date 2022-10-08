package smobilpay

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type requestConfig struct {
	nonce     string
	timestamp time.Time
}

func (config *requestConfig) timestampString() string {
	return fmt.Sprintf("%d", config.timestamp.Unix())
}

func defaultRequestConfig() *requestConfig {
	return &requestConfig{
		nonce:     uuid.NewString(),
		timestamp: time.Now(),
	}
}
