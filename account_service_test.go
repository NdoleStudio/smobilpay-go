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
	assert.Equal(t, 417986.82, account.Balance)
	assert.Equal(t, "XAF", account.Currency)
	assert.Equal(t, "B835273F-0709-AE0C-E01D-B3411C452EEF", account.Key)
	assert.Equal(t, "CM10714-38167", account.AgentID)
	assert.Equal(t, "ACME S3P AGENT", account.AgentName)
	assert.Equal(t, "Adjacent hospital of Buea ", account.AgentAddress)
	assert.Equal(t, "(+237) 673658376", account.AgentPhonenumber)
	assert.Equal(t, "ACME LTD", account.CompanyName)
	assert.Equal(t, "Adjacent Apostolic church Buea,526,Buea", account.CompanyAddress)
	assert.Equal(t, "(+237) 673658376", account.CompanyPhonenumber)
	assert.Equal(t, "10000000.00", account.LimitMax)
	assert.Equal(t, 8950428.00, account.LimitRemaining)

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
