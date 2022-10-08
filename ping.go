package client

import "time"

// PingStatus is the response when pinging the smobilpay API
type PingStatus struct {
	Time    time.Time `json:"time"`
	Version string    `json:"version"`
	Nonce   string    `json:"nonce"`
	Key     string    `json:"key"`
}
