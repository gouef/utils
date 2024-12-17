package tests

import (
	"testing"

	"github.com/gouef/utils"
	"github.com/stretchr/testify/assert"
)

func TestSubstr(t *testing.T) {
	tests := []struct {
		input    string
		start    int
		length   *int
		expected string
	}{
		{"hello world", 0, nil, "hello world"},  // Celý řetězec
		{"hello world", 0, ptr(5), "hello"},     // Začátek + délka
		{"hello world", 6, ptr(5), "world"},     // Od určité pozice
		{"hello world", -5, ptr(3), "wor"},      // Záporný start
		{"hello world", -5, ptr(10), "world"},   // Záporný start přesahující délku
		{"hello world", 11, ptr(5), ""},         // Start mimo délku řetězce
		{"hello world", -12, ptr(5), ""},        // Záporný start mimo délku řetězce
		{"hello world", 3, ptr(20), "lo world"}, // Délka přesahující konec
	}

	for _, tt := range tests {
		result := utils.Substr(tt.input, tt.start, tt.length)
		assert.Equal(t, tt.expected, result, "Input: %s, Start: %d, Length: %v", tt.input, tt.start, tt.length)

		// Test, že Substring funguje stejně
		resultSubstring := utils.Substring(tt.input, tt.start, tt.length)
		assert.Equal(t, tt.expected, resultSubstring, "Substring failed: Input: %s, Start: %d, Length: %v", tt.input, tt.start, tt.length)
	}
}

// Pomocná funkce pro ukazatel na int
func ptr(i int) *int {
	return &i
}

func TestStrpos(t *testing.T) {
	tests := []struct {
		stringStr string
		needle    string
		expected  int
	}{
		{"hello world", "world", 6},        // Výskyt podřetězce
		{"hello world", "hello", 0},        // Výskyt na začátku
		{"hello world hello", "hello", 12}, // Poslední výskyt
		{"hello world", "x", -1},           // Žádný výskyt
		{"hello", "", 5},                   // Prázdný needle
		{"", "hello", -1},                  // Prázdný řetězec
	}

	for _, tt := range tests {
		result := utils.Strpos(tt.stringStr, tt.needle)
		assert.Equal(t, tt.expected, result, "Strpos failed: Input: %s, Needle: %s", tt.stringStr, tt.needle)

		// Test, že StringPosition funguje stejně
		resultPosition := utils.StringPosition(tt.stringStr, tt.needle)
		assert.Equal(t, tt.expected, resultPosition, "StringPosition failed: Input: %s, Needle: %s", tt.stringStr, tt.needle)
	}
}

func TestStrncmp(t *testing.T) {
	tests := []struct {
		s1       string
		s2       string
		n        int
		expected int
	}{
		{"hello", "hello", 5, 0},       // Stejné řetězce
		{"hello", "hella", 5, 1},       // Rozdílné znaky na konci
		{"hello", "hellz", 4, 0},       // Porovnání prvních 4 znaků
		{"hello", "world", 3, -1},      // Rozdílné řetězce
		{"hello", "hello world", 5, 0}, // Prvních 5 znaků je stejné
		{"hello", "he", 5, 1},          // Jeden řetězec kratší
		{"he", "hello", 5, -1},         // Opačný případ
	}

	for _, tt := range tests {
		result := utils.Strncmp(tt.s1, tt.s2, tt.n)
		assert.Equal(t, tt.expected, result, "Strncmp failed: S1: %s, S2: %s, N: %d", tt.s1, tt.s2, tt.n)

		// Test, že StringCompare funguje stejně
		resultCompare := utils.StringCompare(tt.s1, tt.s2, tt.n)
		assert.Equal(t, tt.expected, resultCompare, "StringCompare failed: S1: %s, S2: %s, N: %d", tt.s1, tt.s2, tt.n)
	}
}
