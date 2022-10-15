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
func (service *topupService) GetPackages(ctx context.Context, serviceID string, options ...RequestOption) ([]*Topup, *Response, error) {
	request, err := service.client.newRequest(ctx, options, http.MethodGet, fmt.Sprintf("/topup?serviceid=%s", serviceID), nil)
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	var packages []*Topup
	if err = json.Unmarshal(*response.Body, &packages); err != nil {
		return nil, response, err
	}

	return packages, response, nil
}
