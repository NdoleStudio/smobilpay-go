package smobilpay

// Service represents a service supported by the smobilpay API
type Service struct {
	ServiceID            string     `json:"serviceid"`
	Merchant             string     `json:"merchant"`
	Title                string     `json:"title"`
	Description          string     `json:"description"`
	Category             string     `json:"category"`
	Country              string     `json:"country"`
	LocalCurrency        string     `json:"localCur"`
	Type                 string     `json:"type"`
	Status               string     `json:"status"`
	IsReqCustomerName    int        `json:"isReqCustomerName"`
	IsReqCustomerAddress int        `json:"isReqCustomerAddress"`
	IsReqCustomerNumber  int        `json:"isReqCustomerNumber"`
	IsReqServiceNumber   int        `json:"isReqServiceNumber"`
	LabelCustomerNumber  []I18nText `json:"labelCustomerNumber"`
	LabelServiceNumber   []I18nText `json:"labelServiceNumber"`
	IsVerifiable         bool       `json:"isVerifiable"`
	ValidationMask       *string    `json:"validationMask"`
	Hint                 []I18nText `json:"hint"`
	Denomination         *int       `json:"denomination"`
}

// I18nText represents a localized text entry
type I18nText struct {
	Language  string `json:"language"`
	LocalText string `json:"localText"`
}
