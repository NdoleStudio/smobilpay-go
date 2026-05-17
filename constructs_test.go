package smobilpay

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransaction_IsSuccessful(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		status   string
		expected bool
	}{
		{"returns true for SUCCESS status", "SUCCESS", true},
		{"returns false for PENDING status", "PENDING", false},
		{"returns false for ERRORED status", "ERRORED", false},
		{"returns false for REVERSED status", "REVERSED", false},
		{"returns false for INPROCESS status", "INPROCESS", false},
		{"returns false for empty status", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			transaction := &Transaction{Status: tt.status}
			assert.Equal(t, tt.expected, transaction.IsSuccessful())
		})
	}
}

func TestTransaction_IsReversed(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		status   string
		expected bool
	}{
		{"returns true for REVERSED status", "REVERSED", true},
		{"returns false for SUCCESS status", "SUCCESS", false},
		{"returns false for PENDING status", "PENDING", false},
		{"returns false for ERRORED status", "ERRORED", false},
		{"returns false for INPROCESS status", "INPROCESS", false},
		{"returns false for empty status", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			transaction := &Transaction{Status: tt.status}
			assert.Equal(t, tt.expected, transaction.IsReversed())
		})
	}
}

func TestTransaction_IsFailed(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		status   string
		expected bool
	}{
		{"returns true for ERRORED status", "ERRORED", true},
		{"returns true for ERROREDREFUNDED status", "ERROREDREFUNDED", true},
		{"returns false for SUCCESS status", "SUCCESS", false},
		{"returns false for PENDING status", "PENDING", false},
		{"returns false for REVERSED status", "REVERSED", false},
		{"returns false for empty status", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			transaction := &Transaction{Status: tt.status}
			assert.Equal(t, tt.expected, transaction.IsFailed())
		})
	}
}

func TestTransaction_IsPending(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		status   string
		expected bool
	}{
		{"returns true for PENDING status", "PENDING", true},
		{"returns true for INPROCESS status", "INPROCESS", true},
		{"returns false for SUCCESS status", "SUCCESS", false},
		{"returns false for ERRORED status", "ERRORED", false},
		{"returns false for REVERSED status", "REVERSED", false},
		{"returns false for empty status", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			transaction := &Transaction{Status: tt.status}
			assert.Equal(t, tt.expected, transaction.IsPending())
		})
	}
}
