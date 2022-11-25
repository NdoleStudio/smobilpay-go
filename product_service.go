package smobilpay

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// productService is the API client for the `/product` endpoint
type productService service

// Get cashin pay item
//
// https://apidocs.smobilpay.com/s3papi/API-Reference.2066448558.html
func (service *productService) Get(ctx context.Context, serviceID string, options ...RequestOption) ([]*PayItem, *Response, error) {
	request, err := service.client.newRequest(ctx, options, http.MethodGet, fmt.Sprintf("/product?serviceid=%s", serviceID), nil)
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
