package smobilpay

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWithTimestamp(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	config := defaultRequestConfig()
	timestamp := time.Now().Add(1 * time.Hour)

	// Act
	WithRequestTimestamp(timestamp).apply(config)

	// Assert
	assert.Equal(t, timestamp, config.timestamp)
}

func TestWithNonce(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	config := defaultRequestConfig()
	timestamp := time.Now().Add(1 * time.Hour)

	// Act
	WithRequestTimestamp(timestamp).apply(config)

	// Assert
	assert.Equal(t, timestamp, config.timestamp)
}
