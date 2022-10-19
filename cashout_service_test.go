package smobilpay

import (
	"context"
	"net/http"
	"testing"

	"github.com/NdoleStudio/smobilpay-go/internal/helpers"
	"github.com/NdoleStudio/smobilpay-go/internal/stubs"
	"github.com/stretchr/testify/assert"
)

func TestCashoutService_Get(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.CashoutGetOk())
	accessToken := "6B352110-4716-11ED-963F-0800200C9A66"
	client := New(
		WithBaseURL(server.URL),
		WithAccessToken(accessToken),
		WithAccessSecret("1B875FB0-4717-11ED-963F-0800200C9A66"),
	)
	nonce := "95cdf110-4614-4d95-b6c2-f14fe01c4995"
	serviceID := "20053"

	// Act
	payItems, response, err := client.Cashout.Get(
		context.Background(),
		serviceID,
		WithRequestNonce(nonce),
	)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, 1, len(payItems))
	assert.Equal(t, serviceID, payItems[0].ServiceID)

	// Teardown
	server.Close()
}

func TestCashoutService_GetError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusBadRequest, stubs.CashoutGetError())
	accessToken := "6B352110-4716-11ED-963F-0800200C9A66"
	client := New(
		WithBaseURL(server.URL),
		WithAccessToken(accessToken),
		WithAccessSecret("1B875FB0-4717-11ED-963F-0800200C9A66"),
	)
	nonce := "95cdf110-4614-4d95-b6c2-f14fe01c4995"
	serviceID := "30052"

	// Act
	payItems, response, err := client.Cashout.Get(
		context.Background(),
		serviceID,
		WithRequestNonce(nonce),
	)

	// Assert
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusBadRequest, response.HTTPResponse.StatusCode)
	assert.Equal(t, 0, len(payItems))

	// Teardown
	server.Close()
}
