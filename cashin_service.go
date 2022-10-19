package smobilpay

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// cashinService is the API client for the `/cashin` endpoint
type cashinService service

// Get returns a list of all available cashout packages.
//
// https://apidocs.smobilpay.com/s3papi/API-Reference.2066448558.html
func (service *cashinService) Get(ctx context.Context, serviceID string, options ...RequestOption) ([]*PayItem, *Response, error) {
	request, err := service.client.newRequest(ctx, options, http.MethodGet, fmt.Sprintf("/cashin?serviceid=%s", serviceID), nil)
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	var packages []*PayItem
	if err = json.Unmarshal(*response.Body, &packages); err != nil {
		return nil, response, err
	}

	return packages, response, nil
}
