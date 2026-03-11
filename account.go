package smobilpay

// Account is the response when getting the account information from the smobilpay API
type Account struct {
	Balance            float64 `json:"balance"`
	Currency           string  `json:"currency"`
	Key                string  `json:"key"`
	AgentID            string  `json:"agentId"`
	AgentName          string  `json:"agentName"`
	AgentAddress       string  `json:"agentAddress"`
	AgentPhonenumber   string  `json:"agentPhonenumber"`
	CompanyName        string  `json:"companyName"`
	CompanyAddress     string  `json:"companyAddress"`
	CompanyPhonenumber string  `json:"companyPhonenumber"`
	LimitMax           float64 `json:"limitMax"`
	LimitRemaining     float64 `json:"limitRemaining"`
}
