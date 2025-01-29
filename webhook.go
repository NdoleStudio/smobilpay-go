package smobilpay

// WebhookRequest is the request payload sent to the webhook endpoint
type WebhookRequest struct {
	Timestamp     string `json:"timestamp"`
	TransactionID string `json:"trid"`
	ErrorCode     string `json:"errorCode"`
	Status        string `json:"status"`
}

// IsFailed checks if a WebhookRequest failed
func (request *WebhookRequest) IsFailed() bool {
	return request.Status == "ERRORED" || request.Status == "ERROREDREFUNDED"
}
