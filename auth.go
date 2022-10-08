package smobilpay

import "time"

// AuthParams are additional parameters used for authentication.
type AuthParams struct {
	Nonce     string    `json:"nonce"`
	Timestamp time.Time `json:"timestamp"`
}
