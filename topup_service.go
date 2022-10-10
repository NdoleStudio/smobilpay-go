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
func (service *topupService) GetPackages(ctx context.Context, serviceID string, options ...RequestOption) (*[]TopupPackage, *Response, error) {
	request, err := service.client.newRequest(ctx, options, http.MethodGet, fmt.Sprintf("/topup?serviceid=%s", serviceID), nil)
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	packages := new([]TopupPackage)
	if err = json.Unmarshal(*response.Body, packages); err != nil {
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
func (service *topupService) Collect(ctx context.Context, params *CollectParams, options ...RequestOption) (*map[string]interface{}, *Response, error) {
	request, err := service.client.newRequest(ctx, options, http.MethodPost, "/collectstd", params.toPayload())
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	packages := new(map[string]interface{})
	if err = json.Unmarshal(*response.Body, packages); err != nil {
		return nil, response, err
	}

	return packages, response, nil
}
