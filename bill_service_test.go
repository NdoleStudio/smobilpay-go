package smobilpay

import (
	"context"
	"net/http"
	"testing"

	"github.com/NdoleStudio/smobilpay-go/internal/helpers"
	"github.com/NdoleStudio/smobilpay-go/internal/stubs"
	"github.com/stretchr/testify/assert"
)

func TestBillService_Get(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.BillGet())
	accessToken := "6B352110-4716-11ED-963F-0800200C9A66"
	client := New(
		WithBaseURL(server.URL),
		WithAccessToken(accessToken),
		WithAccessSecret("1B875FB0-4717-11ED-963F-0800200C9A66"),
	)
	nonce := "95cdf110-4614-4d95-b6c2-f14fe01c4995"
	params := &BillGetParams{
		ServiceID:     "10039",
		Merchant:      "ENEO",
		ServiceNumber: "500000001",
	}

	// Act
	bill, response, err := client.Bill.Get(
		context.Background(),
		params,
		WithRequestNonce(nonce),
	)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, 1, len(bill))
	assert.Equal(t, params.ServiceID, bill[0].ServiceID)
	assert.Equal(t, params.Merchant, bill[0].Merchant)
	assert.Equal(t, params.ServiceNumber, bill[0].ServiceNumber)

	// Teardown
	server.Close()
}

func TestBillService_GetEmpty(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.BillEmpty())
	accessToken := "6B352110-4716-11ED-963F-0800200C9A66"
	client := New(
		WithBaseURL(server.URL),
		WithAccessToken(accessToken),
		WithAccessSecret("1B875FB0-4717-11ED-963F-0800200C9A66"),
	)
	nonce := "95cdf110-4614-4d95-b6c2-f14fe01c4995"

	params := &BillGetParams{
		ServiceID:     "10039",
		Merchant:      "ENEO",
		ServiceNumber: "500000001",
	}

	// Act
	bills, _, err := client.Bill.Get(
		context.Background(),
		params,
		WithRequestNonce(nonce),
	)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, 0, len(bills))

	// Teardown
	server.Close()
}
