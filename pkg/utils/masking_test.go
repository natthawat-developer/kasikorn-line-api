package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaskDebitCardNumber(t *testing.T) {
	tests := []struct {
		name           string
		input          *string
		expectedOutput string
	}{
		{
			name:           "Valid debit card number (19 characters including spaces)",
			input:          strPtr("1234 5678 9012 3456"),
			expectedOutput: "1234 56•• •••• 3456",
		},
		{
			name:           "Short card number",
			input:          strPtr("1234 5678"),
			expectedOutput: "", 
		},
		{
			name:           "Nil card number",
			input:          nil,
			expectedOutput: "",
		},
		{
			name:           "Extra long card number",
			input:          strPtr("1234 5678 9012 3456 7890"),
			expectedOutput: "", 
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := MaskDebitCardNumber(tt.input)
			assert.Equal(t, tt.expectedOutput, output, "Unexpected masked debit card number")
		})
	}
}

// Helper function
func strPtr(s string) *string {
	return &s
}
