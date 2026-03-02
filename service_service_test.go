package smobilpay

import (
	"context"
	"net/http"
	"testing"

	"github.com/NdoleStudio/smobilpay-go/internal/helpers"
	"github.com/NdoleStudio/smobilpay-go/internal/stubs"
	"github.com/stretchr/testify/assert"
)

func TestServiceService_GetAll(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.ServiceGetAllOk())
	accessToken := "6B352110-4716-11ED-963F-0800200C9A66"
	client := New(
		WithBaseURL(server.URL),
		WithAccessToken(accessToken),
		WithAccessSecret("1B875FB0-4717-11ED-963F-0800200C9A66"),
	)
	nonce := "95cdf110-4614-4d95-b6c2-f14fe01c4995"

	// Act
	services, response, err := client.Service.GetAll(
		context.Background(),
		WithRequestNonce(nonce),
	)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, 2, len(services))
	assert.Equal(t, "20062", services[0].ServiceID)
	assert.Equal(t, "CMORANGE", services[0].Merchant)
	assert.Equal(t, "ORANGE Recharge/Airtime", services[0].Title)
	assert.Equal(t, "Airtime", services[0].Category)
	assert.Equal(t, "CMR", services[0].Country)
	assert.Equal(t, "XAF", services[0].LocalCurrency)
	assert.Equal(t, "TOPUP", services[0].Type)
	assert.Equal(t, "Active", services[0].Status)
	assert.Equal(t, 0, services[0].IsReqCustomerName)
	assert.Equal(t, 0, services[0].IsReqCustomerAddress)
	assert.Equal(t, 1, services[0].IsReqServiceNumber)
	assert.False(t, services[0].IsVerifiable)
	assert.Equal(t, 50, *services[0].Denomination)
	assert.Nil(t, services[0].LabelCustomerNumber)
	assert.Equal(t, "ORANGE Number", services[0].LabelServiceNumber[0].LocalText)
	assert.Equal(t, "en", services[0].LabelServiceNumber[0].Language)
	assert.Nil(t, services[1].Hint)

	// Teardown
	server.Close()
}

func TestServiceService_GetAllError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusBadRequest, stubs.ServiceGetError())
	accessToken := "6B352110-4716-11ED-963F-0800200C9A66"
	client := New(
		WithBaseURL(server.URL),
		WithAccessToken(accessToken),
		WithAccessSecret("1B875FB0-4717-11ED-963F-0800200C9A66"),
	)
	nonce := "95cdf110-4614-4d95-b6c2-f14fe01c4995"

	// Act
	services, response, err := client.Service.GetAll(
		context.Background(),
		WithRequestNonce(nonce),
	)

	// Assert
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusBadRequest, response.HTTPResponse.StatusCode)
	assert.Equal(t, 0, len(services))

	// Teardown
	server.Close()
}

func TestServiceService_Get(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.ServiceGetOk())
	accessToken := "6B352110-4716-11ED-963F-0800200C9A66"
	client := New(
		WithBaseURL(server.URL),
		WithAccessToken(accessToken),
		WithAccessSecret("1B875FB0-4717-11ED-963F-0800200C9A66"),
	)
	nonce := "95cdf110-4614-4d95-b6c2-f14fe01c4995"
	serviceID := "20062"

	// Act
	svc, response, err := client.Service.Get(
		context.Background(),
		serviceID,
		WithRequestNonce(nonce),
	)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, serviceID, svc.ServiceID)
	assert.Equal(t, "CMORANGE", svc.Merchant)
	assert.Equal(t, "ORANGE Recharge/Airtime", svc.Title)
	assert.Equal(t, "TOPUP", svc.Type)
	assert.Equal(t, 0, svc.IsReqCustomerName)
	assert.False(t, svc.IsVerifiable)
	assert.Equal(t, 50, *svc.Denomination)
	assert.Equal(t, 1, len(svc.Hint))
	assert.Equal(t, "en", svc.Hint[0].Language)
	assert.Nil(t, svc.LabelCustomerNumber)

	// Teardown
	server.Close()
}

func TestServiceService_GetError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusNotFound, stubs.ServiceGetError())
	accessToken := "6B352110-4716-11ED-963F-0800200C9A66"
	client := New(
		WithBaseURL(server.URL),
		WithAccessToken(accessToken),
		WithAccessSecret("1B875FB0-4717-11ED-963F-0800200C9A66"),
	)
	nonce := "95cdf110-4614-4d95-b6c2-f14fe01c4995"
	serviceID := "99999"

	// Act
	svc, response, err := client.Service.Get(
		context.Background(),
		serviceID,
		WithRequestNonce(nonce),
	)

	// Assert
	assert.NotNil(t, err)
	assert.Nil(t, svc)
	assert.Equal(t, http.StatusNotFound, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}
