package tests

import (
	"testing"

	"github.com/gouef/utils"
	"github.com/stretchr/testify/assert"
)

func TestDetectType(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"123", int32(123)},
		{"2147483647", int32(2147483647)},                   // Max int32
		{"9223372036854775807", int64(9223372036854775807)}, // Max int64
		{"3.14", float32(3.14)},
		{"2.718", float32(2.718)},
		{"1.7976931348623157e+308", float64(1.7976931348623157e+308)}, // Max float64
		{"true", true},
		{"false", false},
		{"hello", "hello"},
	}

	for _, tt := range tests {
		result := utils.DetectType(tt.input)
		assert.Equal(t, tt.expected, result, "Input: %s", tt.input)
	}
}

func TestIsInt(t *testing.T) {
	tests := []struct {
		input      string
		expectedOk bool
		expected   int
	}{
		{"123", true, 123},
		{"2147483647", true, 2147483647},  // Max int32
		{"9223372036854775807", false, 0}, // Max int64 (bude vráceno false)
		{"hello", false, 0},
		{"3.14", false, 0},
	}

	for _, tt := range tests {
		ok, result := utils.IsInt(tt.input)
		assert.Equal(t, tt.expectedOk, ok, "Input: %s", tt.input)
		assert.Equal(t, tt.expected, result, "Input: %s", tt.input)
	}
}
