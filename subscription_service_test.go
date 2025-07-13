package smobilpay

import (
	"context"
	"github.com/NdoleStudio/smobilpay-go/internal/helpers"
	"github.com/NdoleStudio/smobilpay-go/internal/stubs"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestSubscriptionService_Get(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.SubscriptionGet())
	accessToken := "6B352110-4716-11ED-963F-0800200C9A66"
	client := New(
		WithBaseURL(server.URL),
		WithAccessToken(accessToken),
		WithAccessSecret("1B875FB0-4717-11ED-963F-0800200C9A66"),
	)
	nonce := "95cdf110-4614-4d95-b6c2-f14fe01c4995"
	params := &SubscriptionGetParams{
		ServiceID:     "10042",
		Merchant:      "CMENEOPREPAID",
		ServiceNumber: "500000001",
	}

	// Act
	subscription, response, err := client.Subscription.Get(context.Background(), params, WithRequestNonce(nonce))

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, 1, len(subscription))
	assert.Equal(t, params.ServiceID, subscription[0].ServiceID)
	assert.Equal(t, params.Merchant, subscription[0].Merchant)
	assert.Equal(t, params.ServiceNumber, subscription[0].ServiceNumber)

	// Teardown
	server.Close()
}
