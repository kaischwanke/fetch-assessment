package utils

import (
	"fetch-assessment/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStripNonAlphanumeric(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "only letters",
			input:    "abcdef",
			expected: "abcdef",
		},
		{
			name:     "only digits",
			input:    "123456",
			expected: "123456",
		},
		{
			name:     "letters and digits",
			input:    "abc123",
			expected: "abc123",
		},
		{
			name:     "letters, digits, and symbols",
			input:    "abc!@#123",
			expected: "abc123",
		},
		{
			name:     "spaces, symbols, and letters",
			input:    " a b @ c # ",
			expected: "abc",
		},
		{
			name:     "mixed letters, digits, spaces, and symbols",
			input:    "a1 b2 c3!@# ",
			expected: "a1b2c3",
		},
		{
			name:     "non-alphanumeric characters only",
			input:    "!@#$%^&*()",
			expected: "",
		},
		{
			name:     "unicode letters",
			input:    "你好世123界",
			expected: "你好世123界",
		},
		{
			name:     "unicode with symbols",
			input:    "こんにちは!!!123",
			expected: "こんにちは123",
		},
		{
			name:     "unicode with spaces and symbols",
			input:    "  مرحبا!123 ",
			expected: "مرحبا123",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := StripNonAlphanumeric(tc.input)
			if result != tc.expected {
				t.Errorf("expected %q, got %q", tc.expected, result)
			}
		})
	}
}

func TestParseTotal(t *testing.T) {
	total, _ := ParseTotal(model.Receipt{Total: "100.00"})
	assert.Equal(t, float64(100), total)
	total, _ = ParseTotal(model.Receipt{Total: "100.11"})
	assert.Equal(t, 100.11, total)
}
