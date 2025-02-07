package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaskAccountNumber(t *testing.T) {
	tests := []struct {
		name           string
		input          *string
		expectedOutput string
	}{
		{
			name:           "Valid account number",
			input:          strPtr("568-2-7661"),
			expectedOutput: "56••••61",
		},
		{
			name:           "Short account number",
			input:          strPtr("12345"),
			expectedOutput: "••••", 
		},
		{
			name:           "Nil account number",
			input:          nil,
			expectedOutput: "••••", 
		},
		{
			name:           "Minimum length (6 chars)",
			input:          strPtr("123456"),
			expectedOutput: "12••••56",
		},
		{
			name:           "Long account number",
			input:          strPtr("9876543210987654"),
			expectedOutput: "98••••54",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := MaskAccountNumber(tt.input)
			assert.NotNil(t, output, "Output should not be nil")
			assert.Equal(t, tt.expectedOutput, *output, "Unexpected masked account number")
		})
	}
}

// strPtr แปลง string เป็น *string (ช่วยให้เขียนเทสต์ง่ายขึ้น)
func strPtr(s string) *string {
	return &s
}
