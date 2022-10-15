package smobilpay

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// topupService is the API client for the `/` endpoint
type topupService service

// GetPackages returns a list of all available topup packages.
//
// https://apidocs.smobilpay.com/s3papi/API-Reference.2066448558.html
func (service *topupService) GetPackages(ctx context.Context, serviceID string, options ...RequestOption) ([]*TopupPackage, *Response, error) {
	request, err := service.client.newRequest(ctx, options, http.MethodGet, fmt.Sprintf("/topup?serviceid=%s", serviceID), nil)
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	var packages []*TopupPackage
	if err = json.Unmarshal(*response.Body, &packages); err != nil {
		return nil, response, err
	}

	return packages, response, nil
}

// Quote initializes the airtime topup transaction
//
// https://apidocs.smobilpay.com/s3papi/API-Reference.2066448558.html
func (service *topupService) Quote(ctx context.Context, params *QuoteParams, options ...RequestOption) (*Quote, *Response, error) {
	request, err := service.client.newRequest(ctx, options, http.MethodPost, "/quotestd", map[string]string{
		"payItemId": params.PayItemID,
		"amount":    params.Amount,
	})
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	packages := new(Quote)
	if err = json.Unmarshal(*response.Body, packages); err != nil {
		return nil, response, err
	}

	return packages, response, nil
}

// Collect confirms the airtime topup transaction
//
// https://apidocs.smobilpay.com/s3papi/API-Reference.2066448558.html
func (service *topupService) Collect(ctx context.Context, params *CollectParams, options ...RequestOption) (*Transaction, *Response, error) {
	request, err := service.client.newRequest(ctx, options, http.MethodPost, "/collectstd", params.toPayload())
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	transaction := new(Transaction)
	if err = json.Unmarshal(*response.Body, transaction); err != nil {
		return nil, response, err
	}

	return transaction, response, nil
}

// Verify gets the current payment collection status
//
// https://apidocs.smobilpay.com/s3papi/API-Reference.2066448558.html
func (service *topupService) Verify(ctx context.Context, paymentTransactionNumber string, options ...RequestOption) (*Transaction, *Response, error) {
	request, err := service.client.newRequest(ctx, options, http.MethodGet, fmt.Sprintf("/verifytx?ptn=%s", paymentTransactionNumber), nil)
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	var transactions []Transaction
	if err = json.Unmarshal(*response.Body, &transactions); err != nil {
		return nil, response, err
	}

	if len(transactions) == 0 {
		return nil, response, fmt.Errorf("cannot verify transaction with payment transaction number [%s]", paymentTransactionNumber)
	}

	return &transactions[0], response, nil
}
