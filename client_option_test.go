package smobilpay

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWithHTTPClient(t *testing.T) {
	t.Run("httpClient is not set when the httpClient is nil", func(t *testing.T) {
		// Setup
		t.Parallel()

		// Arrange
		config := defaultClientConfig()

		// Act
		WithHTTPClient(nil).apply(config)

		// Assert
		assert.NotNil(t, config.httpClient)
	})

	t.Run("httpClient is set when the httpClient is not nil", func(t *testing.T) {
		// Setup
		t.Parallel()

		// Arrange
		config := defaultClientConfig()
		newClient := &http.Client{Timeout: 300}

		// Act
		WithHTTPClient(newClient).apply(config)

		// Assert
		assert.NotNil(t, config.httpClient)
		assert.Equal(t, newClient.Timeout, config.httpClient.Timeout)
	})
}

func TestWithBaseURL(t *testing.T) {
	t.Run("baseURL is set successfully", func(t *testing.T) {
		// Setup
		t.Parallel()

		// Arrange
		baseURL := "https://example.com"
		config := defaultClientConfig()

		// Act
		WithBaseURL(baseURL).apply(config)

		// Assert
		assert.Equal(t, config.baseURL, config.baseURL)
	})

	t.Run("tailing / is trimmed from baseURL", func(t *testing.T) {
		// Setup
		t.Parallel()

		// Arrange
		baseURL := "https://example.com/"
		config := defaultClientConfig()

		// Act
		WithBaseURL(baseURL).apply(config)

		// Assert
		assert.Equal(t, "https://example.com", config.baseURL)
	})
}

func TestWithAccessToken(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	config := defaultClientConfig()
	accessToken := "access-token"

	// Act
	WithAccessToken(accessToken).apply(config)

	// Assert
	assert.Equal(t, accessToken, config.accessToken)
}

func TestWithAccessSecret(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	config := defaultClientConfig()
	accessSecret := "access-secret"

	// Act
	WithAccessSecret(accessSecret).apply(config)

	// Assert
	assert.Equal(t, accessSecret, config.accessSecret)
}

func TestWithCollectSyncVerifyInterval(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	config := defaultClientConfig()
	interval := time.Second * 2

	// Act
	WithCollectSyncVerifyInterval(interval).apply(config)

	// Assert
	assert.Equal(t, interval, config.collectSyncVerifyInterval)
}

func TestWithCollectSyncVerifyRetryCount(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	config := defaultClientConfig()
	retryCount := uint(1000)

	// Act
	WithCollectSyncVerifyRetryCount(retryCount).apply(config)

	// Assert
	assert.Equal(t, retryCount, config.collectSyncVerifyRetryCount)
}
