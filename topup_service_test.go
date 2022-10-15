package smobilpay

import (
	"context"
	"net/http"
	"testing"

	"github.com/NdoleStudio/smobilpay-go/internal/helpers"
	"github.com/NdoleStudio/smobilpay-go/internal/stubs"
	"github.com/stretchr/testify/assert"
)

func TestTopupService_GetPackages(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.TopupGetPackagesOk())
	accessToken := "6B352110-4716-11ED-963F-0800200C9A66"
	client := New(
		WithBaseURL(server.URL),
		WithAccessToken(accessToken),
		WithAccessSecret("1B875FB0-4717-11ED-963F-0800200C9A66"),
	)
	nonce := "95cdf110-4614-4d95-b6c2-f14fe01c4995"

	// Act
	topupPackages, response, err := client.Topup.GetPackages(
		context.Background(),
		"1234",
		WithRequestNonce(nonce),
	)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, 2, len(topupPackages))

	// Teardown
	server.Close()
}

func TestTopupService_Quote(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.TopupQuoteOk())
	accessToken := "6B352110-4716-11ED-963F-0800200C9A66"
	client := New(
		WithBaseURL(server.URL),
		WithAccessToken(accessToken),
		WithAccessSecret("1B875FB0-4717-11ED-963F-0800200C9A66"),
	)
	nonce := "95cdf110-4614-4d95-b6c2-f14fe01c4995"
	payItemID := "S-112-951-CMORANGE-20062-CM_ORANGE_VTU_CUSTOM-1"

	// Act
	quote, response, err := client.Topup.Quote(
		context.Background(),
		&QuoteParams{
			PayItemID: payItemID,
			Amount:    "100",
		},
		WithRequestNonce(nonce),
	)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, payItemID, quote.PayItemID)
	assert.Equal(t, 100, quote.PriceSystemCurrency)

	// Teardown
	server.Close()
}

func TestTopupService_Collect(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.TopupCollectOk())
	accessToken := "6B352110-4716-11ED-963F-0800200C9A66"
	client := New(
		WithBaseURL(server.URL),
		WithAccessToken(accessToken),
		WithAccessSecret("1B875FB0-4717-11ED-963F-0800200C9A66"),
	)
	nonce := "95cdf110-4614-4d95-b6c2-f14fe01c4995"
	params := &CollectParams{
		QuoteID:               "15380e55-6227-4a25-8e1f-23c8735ce242",
		CustomerPhoneNumber:   "697777777",
		CustomerEmailAddress:  "dev@test.com",
		ServiceNumber:         "697777777",
		ExternalTransactionID: "999992624813740205",
	}

	// Act
	transaction, response, err := client.Topup.Collect(
		context.Background(),
		params,
		WithRequestNonce(nonce),
	)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, params.ExternalTransactionID, *transaction.ExternalTransactionID)
	assert.False(t, transaction.IsFailed())
	assert.Equal(t, "SUCCESS", transaction.Status)

	// Teardown
	server.Close()
}

func TestTopupService_Verify(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.TopupVerifyOk())
	accessToken := "6B352110-4716-11ED-963F-0800200C9A66"
	client := New(
		WithBaseURL(server.URL),
		WithAccessToken(accessToken),
		WithAccessSecret("1B875FB0-4717-11ED-963F-0800200C9A66"),
	)
	nonce := "95cdf110-4614-4d95-b6c2-f14fe01c4995"
	paymentTransactionNumber := "99999166542651400095315364801168"

	// Act
	transaction, response, err := client.Topup.Verify(
		context.Background(),
		paymentTransactionNumber,
		WithRequestNonce(nonce),
	)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, paymentTransactionNumber, transaction.PaymentTransactionNumber)
	assert.False(t, transaction.IsFailed())
	assert.Equal(t, "SUCCESS", transaction.Status)

	// Teardown
	server.Close()
}

func TestTopupService_VerifyEmpty(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.TopupVerifyEmpty())
	accessToken := "6B352110-4716-11ED-963F-0800200C9A66"
	client := New(
		WithBaseURL(server.URL),
		WithAccessToken(accessToken),
		WithAccessSecret("1B875FB0-4717-11ED-963F-0800200C9A66"),
	)
	nonce := "95cdf110-4614-4d95-b6c2-f14fe01c4995"
	paymentTransactionNumber := "99999166542651400095315364801168"

	// Act
	_, _, err := client.Topup.Verify(
		context.Background(),
		paymentTransactionNumber,
		WithRequestNonce(nonce),
	)

	// Assert
	assert.NotNil(t, err)

	// Teardown
	server.Close()
}
