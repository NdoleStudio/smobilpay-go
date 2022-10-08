package client

import (
	"context"
	"encoding/json"
	"net/http"
)

// topupService is the API client for the `/` endpoint
type topupService service

// Get topup items
func (service *topupService) Get(ctx context.Context, serviceID string) (*[]map[string]interface{}, *Response, error) {
	request, err := service.client.newRequest(ctx, http.MethodGet, "/200", nil)
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	status := new([]map[string]interface{})
	if err = json.Unmarshal(*response.Body, status); err != nil {
		return nil, response, err
	}

	return status, response, nil
}
