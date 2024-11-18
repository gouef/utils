package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRound(t *testing.T) {
	result := RoundTo(1234, 1)
	assert.Equal(t, float64(1230), result, "Expected 1230")

	result1 := RoundTo(1234, 2)
	assert.Equal(t, float64(1200), result1, "Expected 1200")

	result2 := RoundTo(1234, 3)
	assert.Equal(t, float64(1000), result2, "Expected 1000")
}

func TestRoundTens(t *testing.T) {
	result := RoundTens(123)
	assert.Equal(t, float64(120), result, "Expected 120")
}
