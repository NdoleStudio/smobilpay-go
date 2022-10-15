package smobilpay

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// billService is the API client for the `/bill` endpoint
type billService service

const (
	// BillMerchantENEO is the merchant name for ENEO
	BillMerchantENEO = "ENEO"

	// BillMerchantCamwater is the merchant name for camwater
	BillMerchantCamwater = "CAMWATER"
)

// Get returns a bill
//
// https://apidocs.smobilpay.com/s3papi/API-Reference.2066448558.html
func (service *billService) Get(ctx context.Context, params *BillGetParams, options ...RequestOption) ([]*Bill, *Response, error) {
	request, err := service.client.newRequest(ctx, options, http.MethodGet, fmt.Sprintf("/bill?serviceid=%s&merchant=%s&serviceNumber=%s", params.ServiceID, params.Merchant, params.ServiceNumber), nil)
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	var bills []*Bill
	if err = json.Unmarshal(*response.Body, &bills); err != nil {
		return nil, response, err
	}

	return bills, response, nil
}
