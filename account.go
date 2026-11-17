package smobilpay

import (
	"encoding/json"
	"fmt"
)

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
	LimitMax           string  `json:"limitMax"`
	LimitRemaining     float64 `json:"limitRemaining"`
}

// UnmarshalJSON accepts balances encoded as numbers or numeric strings.
func (account *Account) UnmarshalJSON(data []byte) error {
	type accountJSON Account
	payload := struct {
		*accountJSON
		Balance json.Number `json:"balance"`
	}{
		accountJSON: (*accountJSON)(account),
	}
	if err := json.Unmarshal(data, &payload); err != nil {
		return err
	}
	if payload.Balance == "" {
		return nil
	}
	balance, err := payload.Balance.Float64()
	if err != nil {
		return fmt.Errorf("cannot parse account balance: %w", err)
	}
	account.Balance = balance
	return nil
}
