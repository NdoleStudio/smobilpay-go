package smobilpay

// SubscriptionGetParams are the parameters for getting a bill
type SubscriptionGetParams struct {
	ServiceID      string
	Merchant       string
	ServiceNumber  string
	CustomerNumber string
}

// Subscription represents a service offered by the smobilpay API
type Subscription struct {
	ServiceNumber      string  `json:"serviceNumber"`
	ServiceID          string  `json:"serviceid"`
	Merchant           string  `json:"merchant"`
	PayItemID          string  `json:"payItemId"`
	PayItemDescription string  `json:"payItemDescr"`
	AmountType         string  `json:"amountType"`
	Name               string  `json:"name"`
	LocalCur           string  `json:"localCur"`
	AmountLocalCur     string  `json:"amountLocalCur"`
	CustomerReference  string  `json:"customerReference"`
	CustomerName       string  `json:"customerName"`
	CustomerNumber     string  `json:"customerNumber"`
	StartDate          string  `json:"startDate"`
	DueDate            string  `json:"dueDate"`
	EndDate            string  `json:"endDate"`
	OptionalString     *string `json:"optStrg"`
	OptionalNumber     *int    `json:"optNmb"`
}
