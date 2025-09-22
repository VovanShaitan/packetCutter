package services

import (
	"packetCutter/internal/domain"
	"packetCutter/internal/storage"
	"testing"
)

func TestMatcherService_CountMatchesForSequence(t *testing.T) {
	matcher := NewMatcherService(storage.NewResultCollection())

	tests := []struct {
		name          string
		target        string
		sequence      domain.Prediction
		expected      int
		expectedError error
	}{
		{
			name:          "valid sequence",
			target:        "123",
			sequence:      domain.Prediction{"1", "2", "3"},
			expected:      3,
			expectedError: nil,
		},
		{
			name:          "sequence length mismatch",
			target:        "123",
			sequence:      domain.Prediction{"1", "2"},
			expected:      0,
			expectedError: domain.ErrPredictLength,
		},
		{
			name:          "multi char sequence",
			target:        "123",
			sequence:      domain.Prediction{"12", "23", "34"},
			expected:      3,
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			count, err := matcher.CountMatchesForPrediction(tt.target, tt.sequence)

			if count != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, count)
			}

			if (err != nil) != (tt.expectedError != nil) {
				t.Errorf("error: expected %v, got %v", tt.expectedError, err)
			}
		})
	}
}
