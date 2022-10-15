package smobilpay

import (
	"time"
)

// QuoteParams is the input needed to initialize a transaction
type QuoteParams struct {
	PayItemID string `json:"payItemId"`
	Amount    string `json:"amount"`
}

// Quote represents an initialized transaction
type Quote struct {
	QuoteID             string      `json:"quoteId"`
	ExpiresAt           time.Time   `json:"expiresAt"`
	PayItemID           string      `json:"payItemId"`
	AmountLocalCurrency string      `json:"amountLocalCur"`
	PriceLocalCurrency  int         `json:"priceLocalCur"`
	PriceSystemCurrency int         `json:"priceSystemCur"`
	LocalCurrency       string      `json:"localCur"`
	SystemCurrency      string      `json:"systemCur"`
	Promotion           interface{} `json:"promotion"`
}

// CollectParams is the input needed to confirm a transaction
type CollectParams struct {
	QuoteID               string `json:"quoteId"`
	CustomerPhoneNumber   string `json:"customerPhonenumber"`
	CustomerEmailAddress  string `json:"customerEmailaddress"`
	CustomerName          string `json:"customerName"`
	CustomerAddress       string `json:"customerAddress"`
	CustomerNumber        string `json:"customerNumber"`
	ServiceNumber         string `json:"serviceNumber"`
	ExternalTransactionID string `json:"trid"`
}

func (params *CollectParams) toPayload() map[string]string {
	payload := map[string]string{
		"quoteId":              params.QuoteID,
		"customerPhonenumber":  params.CustomerPhoneNumber,
		"customerEmailaddress": params.CustomerEmailAddress,
	}

	if params.CustomerName != "" {
		payload["customerName"] = params.CustomerName
	}

	if params.CustomerAddress != "" {
		payload["customerAddress"] = params.CustomerAddress
	}

	if params.ServiceNumber != "" {
		payload["serviceNumber"] = params.ServiceNumber
	}

	if params.ExternalTransactionID != "" {
		payload["trid"] = params.ExternalTransactionID
	}

	return payload
}

// Transaction represents a transaction
type Transaction struct {
	PaymentTransactionNumber string      `json:"ptn"`
	Timestamp                time.Time   `json:"timestamp"`
	AgentBalance             string      `json:"agentBalance"`
	ReceiptNumber            string      `json:"receiptNumber"`
	VerificationCode         string      `json:"veriCode"`
	PriceLocalCurrency       string      `json:"priceLocalCur"`
	PriceSystemCurrency      string      `json:"priceSystemCur"`
	LocalCurrency            string      `json:"localCur"`
	SystemCurrency           string      `json:"systemCur"`
	ExternalTransactionID    *string     `json:"trid"`
	Pin                      interface{} `json:"pin"`
	Status                   string      `json:"status"`
	PayItemDescription       *string     `json:"payItemDescr"`
	PayItemID                string      `json:"payItemId"`
}

// IsFailed checks if a transaction failed
func (transaction *Transaction) IsFailed() bool {
	return transaction.Status == "ERRORED" || transaction.Status == "ERROREDREFUNDED"
}
