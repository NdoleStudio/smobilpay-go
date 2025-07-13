package smobilpay

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// subscriptionService is the API client for the `/subscription` endpoint
type subscriptionService service

// Get subscription payment handler
//
// https://apidocs.smobilpay.com/s3papi/API-Reference.2066448558.html#APIReference-Specification
func (service *subscriptionService) Get(ctx context.Context, params *SubscriptionGetParams, options ...RequestOption) ([]*Subscription, *Response, error) {
	q := url.Values{
		"merchant":  []string{params.Merchant},
		"serviceid": []string{params.ServiceID},
	}
	if params.ServiceNumber != "" {
		q.Set("serviceNumber", params.ServiceNumber)
	}
	if params.CustomerNumber != "" {
		q.Set("customerNumber", params.CustomerNumber)
	}
	request, err := service.client.newRequest(ctx, options, http.MethodGet, fmt.Sprintf("/subscription?%s", q.Encode()), nil)
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	var services []*Subscription
	if err = json.Unmarshal(*response.Body, &services); err != nil {
		return nil, response, err
	}

	return services, response, nil
}
