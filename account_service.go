package smobilpay

import (
	"context"
	"encoding/json"
	"net/http"
)

// accountService is the API client for the `/account` endpoint
type accountService service

// Get returns the account information and remaining balance
//
// https://apidocs.smobilpay.com/s3papi/API-Reference.2066448558.html
func (service *accountService) Get(ctx context.Context, options ...RequestOption) (*Account, *Response, error) {
	request, err := service.client.newRequest(ctx, options, http.MethodGet, "/account", nil)
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	account := new(Account)
	if err = json.Unmarshal(*response.Body, account); err != nil {
		return nil, response, err
	}

	return account, response, nil
}
