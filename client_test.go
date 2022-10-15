package smobilpay

import (
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/NdoleStudio/smobilpay-go/internal/helpers"
	"github.com/NdoleStudio/smobilpay-go/internal/stubs"
	"github.com/stretchr/testify/assert"
)

func TestClient_PingOk(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.PingOk())
	accessToken := "6B352110-4716-11ED-963F-0800200C9A66"
	client := New(
		WithBaseURL(server.URL),
		WithAccessToken(accessToken),
		WithAccessSecret("1B875FB0-4717-11ED-963F-0800200C9A66"),
	)
	nonce := "95cdf110-4614-4d95-b6c2-f14fe01c4995"

	// Act
	status, response, err := client.Ping(
		context.Background(),
		WithRequestTimestamp(time.Now()),
		WithRequestNonce(nonce),
	)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, nonce, status.Nonce)
	assert.Equal(t, accessToken, status.Key)

	// Teardown
	server.Close()
}

func TestClient_PingError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusBadRequest, stubs.PingError())
	accessToken := "6B352110-4716-11ED-963F-0800200C9A66"
	client := New(
		WithBaseURL(server.URL),
		WithAccessToken(accessToken),
		WithAccessSecret("1B875FB0-4717-11ED-963F-0800200C9A66"),
	)
	nonce := "95cdf110-4614-4d95-b6c2-f14fe01c4995"

	// Act
	status, response, err := client.Ping(
		context.Background(),
		WithRequestTimestamp(time.Now()),
		WithRequestNonce(nonce),
	)

	// Assert
	assert.NotNil(t, err)
	assert.Nil(t, status)
	assert.Equal(t, http.StatusBadRequest, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}

func TestClient_PingRequest(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	request := new(http.Request)
	server := helpers.MakeRequestCapturingTestServer(http.StatusOK, stubs.PingOk(), request)
	accessToken := "6B352110-4716-11ED-963F-0800200C9A66"
	accessSecret := "1B875FB0-4717-11ED-963F-0800200C9A66"
	client := New(
		WithBaseURL(server.URL),
		WithAccessToken(accessToken),
		WithAccessSecret(accessSecret),
	)
	nonce := "95cdf110-4614-4d95-b6c2-f14fe01c4995"
	signature := computeHmac(
		"GET&"+url.QueryEscape(server.URL)+"%2Fping&s3pAuth_nonce%3D95cdf110-4614-4d95-b6c2-f14fe01c4995%26s3pAuth_signature_method%3DHMAC-SHA1%26s3pAuth_timestamp%3D1613869830%26s3pAuth_token%3D6B352110-4716-11ED-963F-0800200C9A66",
		accessSecret,
	)

	// Act
	_, _, err := client.Ping(
		context.Background(),
		WithRequestTimestamp(time.Date(2021, time.Month(2), 21, 1, 10, 30, 0, time.UTC)),
		WithRequestNonce(nonce),
	)

	// Assert
	assert.Nil(t, err)
	assert.Equal(
		t,
		"s3pAuth,s3pAuth_nonce=\"95cdf110-4614-4d95-b6c2-f14fe01c4995\",s3pAuth_signature=\""+signature+"\",s3pAuth_signature_method=\"HMAC-SHA1\",s3pAuth_timestamp=\"1613869830\",s3pAuth_token=\"6B352110-4716-11ED-963F-0800200C9A66\"",
		request.Header.Get("Authorization"),
	)

	// Teardown
	server.Close()
}

func TestClient_Quote(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.QuoteOk())
	accessToken := "6B352110-4716-11ED-963F-0800200C9A66"
	client := New(
		WithBaseURL(server.URL),
		WithAccessToken(accessToken),
		WithAccessSecret("1B875FB0-4717-11ED-963F-0800200C9A66"),
	)
	nonce := "95cdf110-4614-4d95-b6c2-f14fe01c4995"
	payItemID := "S-112-951-CMORANGE-20062-CM_ORANGE_VTU_CUSTOM-1"

	// Act
	quote, response, err := client.Quote(
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

func TestClient_Collect(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.CollectOk())
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
	transaction, response, err := client.Collect(
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

func TestClient_Verify(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.VerifyOk())
	accessToken := "6B352110-4716-11ED-963F-0800200C9A66"
	client := New(
		WithBaseURL(server.URL),
		WithAccessToken(accessToken),
		WithAccessSecret("1B875FB0-4717-11ED-963F-0800200C9A66"),
	)
	nonce := "95cdf110-4614-4d95-b6c2-f14fe01c4995"
	paymentTransactionNumber := "99999166542651400095315364801168"

	// Act
	transaction, response, err := client.Verify(
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

func TestClient_TransactionHistory(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.VerifyOk())
	accessToken := "6B352110-4716-11ED-963F-0800200C9A66"
	client := New(
		WithBaseURL(server.URL),
		WithAccessToken(accessToken),
		WithAccessSecret("1B875FB0-4717-11ED-963F-0800200C9A66"),
	)
	paymentTransactionNumber := "99999166542651400095315364801168"

	// Act
	transactions, response, err := client.TransactionHistory(context.Background(), time.Now(), time.Now())

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, 1, len(transactions))
	assert.Equal(t, paymentTransactionNumber, transactions[0].PaymentTransactionNumber)
	assert.False(t, transactions[0].IsFailed())
	assert.Equal(t, "SUCCESS", transactions[0].Status)

	// Teardown
	server.Close()
}

func TestClient_VerifyEmpty(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.VerifyEmpty())
	accessToken := "6B352110-4716-11ED-963F-0800200C9A66"
	client := New(
		WithBaseURL(server.URL),
		WithAccessToken(accessToken),
		WithAccessSecret("1B875FB0-4717-11ED-963F-0800200C9A66"),
	)
	nonce := "95cdf110-4614-4d95-b6c2-f14fe01c4995"
	paymentTransactionNumber := "99999166542651400095315364801168"

	// Act
	_, _, err := client.Verify(
		context.Background(),
		paymentTransactionNumber,
		WithRequestNonce(nonce),
	)

	// Assert
	assert.NotNil(t, err)

	// Teardown
	server.Close()
}

func computeHmac(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha1.New, key)
	_, _ = h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
