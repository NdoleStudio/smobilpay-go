package smobilpay

import "time"

// BillGetParams are the parameters for getting a bill
type BillGetParams struct {
	ServiceID     string
	Merchant      string
	ServiceNumber string
}

// Bill is the details for a bill payment
type Bill struct {
	PayItemID           string      `json:"payItemId"`
	ServiceNumber       string      `json:"serviceNumber"`
	ServiceID           string      `json:"serviceid"`
	Merchant            string      `json:"merchant"`
	AmountType          string      `json:"amountType"`
	LocalCurrency       string      `json:"localCur"`
	AmountLocalCurrency string      `json:"amountLocalCur"`
	BillNumber          string      `json:"billNumber"`
	CustomerNumber      string      `json:"customerNumber"`
	BillMonth           string      `json:"billMonth"`
	BillYear            string      `json:"billYear"`
	BillDate            time.Time   `json:"billDate"`
	BillDueDate         time.Time   `json:"billDueDate"`
	PayItemDescription  string      `json:"payItemDescr"`
	BillType            string      `json:"billType"`
	PenaltyAmount       string      `json:"penaltyAmount"`
	PayOrder            int         `json:"payOrder"`
	OptionalString      interface{} `json:"optStrg"`
	OptionalNumber      interface{} `json:"optNmb"`
}
