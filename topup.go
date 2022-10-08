package smobilpay

// TopupPackage represents a network where we can buy airtime credit
type TopupPackage struct {
	ServiceID           string      `json:"serviceid"`
	Merchant            string      `json:"merchant"`
	PayItemID           string      `json:"payItemId"`
	AmountType          string      `json:"amountType"`
	LocalCurrency       string      `json:"localCur"`
	Name                string      `json:"name"`
	AmountLocalCurrency interface{} `json:"amountLocalCur"`
	Description         string      `json:"description"`
	PayItemDescription  interface{} `json:"payItemDescr"`
	OptionalString      interface{} `json:"optStrg"`
	OptionalNumber      interface{} `json:"optNmb"`
}
