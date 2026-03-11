package smobilpay

import (
	"context"
	"net/http"
	"testing"

	"github.com/NdoleStudio/smobilpay-go/internal/helpers"
	"github.com/NdoleStudio/smobilpay-go/internal/stubs"
	"github.com/stretchr/testify/assert"
)

func TestAccountService_Get(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.AccountGet())
	accessToken := "6B352110-4716-11ED-963F-0800200C9A66"
	client := New(
		WithBaseURL(server.URL),
		WithAccessToken(accessToken),
		WithAccessSecret("1B875FB0-4717-11ED-963F-0800200C9A66"),
	)
	nonce := "95cdf110-4614-4d95-b6c2-f14fe01c4995"

	// Act
	account, response, err := client.Account.Get(
		context.Background(),
		WithRequestNonce(nonce),
	)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, float64(234500.00), account.Balance)
	assert.Equal(t, "XAF", account.Currency)
	assert.Equal(t, accessToken, account.Key)
	assert.Equal(t, "AGT-001-00234", account.AgentID)
	assert.Equal(t, "John Doe", account.AgentName)
	assert.Equal(t, "123 Main Street, Douala", account.AgentAddress)
	assert.Equal(t, "237699999999", account.AgentPhonenumber)
	assert.Equal(t, "Acme Payments Ltd", account.CompanyName)
	assert.Equal(t, "456 Business Ave, Yaounde", account.CompanyAddress)
	assert.Equal(t, "237677777777", account.CompanyPhonenumber)
	assert.Equal(t, float64(5000000.00), account.LimitMax)
	assert.Equal(t, float64(4765500.00), account.LimitRemaining)

	// Teardown
	server.Close()
}

func TestAccountService_GetError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusUnauthorized, stubs.AccountGetError())
	accessToken := "6B352110-4716-11ED-963F-0800200C9A66"
	client := New(
		WithBaseURL(server.URL),
		WithAccessToken(accessToken),
		WithAccessSecret("1B875FB0-4717-11ED-963F-0800200C9A66"),
	)
	nonce := "95cdf110-4614-4d95-b6c2-f14fe01c4995"

	// Act
	account, response, err := client.Account.Get(
		context.Background(),
		WithRequestNonce(nonce),
	)

	// Assert
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusUnauthorized, response.HTTPResponse.StatusCode)
	assert.Nil(t, account)

	// Teardown
	server.Close()
}
