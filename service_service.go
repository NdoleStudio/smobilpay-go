package smobilpay

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// serviceService is the API client for the `/service` endpoint
type serviceService service

// GetAll retrieves the list of services supported by the system.
//
// https://apidocs.smobilpay.com/s3papi/API-Reference.2066448558.html
func (service *serviceService) GetAll(ctx context.Context, options ...RequestOption) ([]*Service, *Response, error) {
	request, err := service.client.newRequest(ctx, options, http.MethodGet, "/service", nil)
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	var services []*Service
	fmt.Printf("%s", string(*response.Body))
	if err = json.Unmarshal(*response.Body, &services); err != nil {
		return nil, response, err
	}

	return services, response, nil
}

// Get retrieves a single service by its ID.
//
// https://apidocs.smobilpay.com/s3papi/API-Reference.2066448558.html
func (service *serviceService) Get(ctx context.Context, serviceID string, options ...RequestOption) (*Service, *Response, error) {
	request, err := service.client.newRequest(ctx, options, http.MethodGet, fmt.Sprintf("/service/%s", serviceID), nil)
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	result := new(Service)
	if err = json.Unmarshal(*response.Body, result); err != nil {
		return nil, response, err
	}

	return result, response, nil
}
