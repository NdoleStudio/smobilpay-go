package smobilpay

import (
	"time"

	"github.com/davecgh/go-spew/spew"
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
	QuoteID               string  `json:"quoteId"`
	CustomerPhoneNumber   string  `json:"customerPhonenumber"`
	CustomerEmailAddress  string  `json:"customerEmailaddress"`
	CustomerName          *string `json:"customerName"`
	CustomerAddress       *string `json:"customerAddress"`
	CustomerNumber        *string `json:"customerNumber"`
	ServiceNumber         *string `json:"serviceNumber"`
	ExternalTransactionID *string `json:"trid"`
}

func (params *CollectParams) toPayload() map[string]string {
	payload := map[string]string{
		"quoteId":              params.QuoteID,
		"customerPhonenumber":  params.CustomerPhoneNumber,
		"customerEmailaddress": params.CustomerEmailAddress,
	}

	if params.CustomerName != nil {
		payload["customerName"] = *params.CustomerName
	}

	if params.CustomerAddress != nil {
		payload["customerAddress"] = *params.CustomerAddress
	}

	if params.ServiceNumber != nil {
		payload["serviceNumber"] = *params.ServiceNumber
	}

	if params.ExternalTransactionID != nil {
		payload["trid"] = *params.ExternalTransactionID
	}

	spew.Dump(payload)

	return payload
}
