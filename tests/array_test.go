package tests

import (
	"fmt"
	"github.com/gouef/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInArray(t *testing.T) {
	tests := []struct {
		needle   string
		haystack []string
		expected bool
	}{
		{"apple", []string{"apple", "banana", "cherry"}, true},
		{"orange", []string{"apple", "banana", "cherry"}, false},
		{"", []string{"apple", "banana", "cherry"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.needle, func(t *testing.T) {
			result := utils.InArray(tt.needle, tt.haystack)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestInListArray(t *testing.T) {
	tests := []struct {
		needle   []string
		haystack []string
		expected bool
	}{
		{[]string{"apple"}, []string{"apple", "banana", "cherry"}, true},
		{[]string{"cherry"}, []string{"apple", "banana", "cherry"}, true},
		{[]string{"pineapple", "banana"}, []string{"apple", "banana", "cherry", "pineapple"}, true},
		{[]string{"orange"}, []string{"apple", "banana", "cherry"}, false},
		{[]string{""}, []string{"apple", "banana", "cherry"}, false},
	}

	for _, tt := range tests {
		t.Run(utils.Implode("-", tt.needle), func(t *testing.T) {
			result := utils.InListArray(tt.needle, tt.haystack)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestExplode(t *testing.T) {
	tests := []struct {
		separator string
		input     string
		expected  []string
	}{
		{",", "apple,banana,cherry", []string{"apple", "banana", "cherry"}},
		{" ", "apple banana cherry", []string{"apple", "banana", "cherry"}},
		{";", "apple;banana;cherry", []string{"apple", "banana", "cherry"}},
		{"", "apple", []string{"apple"}},
	}

	for _, tt := range tests {
		t.Run(tt.separator, func(t *testing.T) {
			result := utils.Explode(tt.separator, tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestImplode(t *testing.T) {
	tests := []struct {
		separator string
		elements  []string
		expected  string
	}{
		{",", []string{"apple", "banana", "cherry"}, "apple,banana,cherry"},
		{" ", []string{"apple", "banana", "cherry"}, "apple banana cherry"},
		{";", []string{"apple", "banana", "cherry"}, "apple;banana;cherry"},
		{"", []string{"apple"}, "apple"},
	}

	for _, tt := range tests {
		t.Run(tt.separator, func(t *testing.T) {
			result := utils.Implode(tt.separator, tt.elements)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestSlice(t *testing.T) {
	tests := []struct {
		arr      []int
		start    int
		length   int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, 1, 3, []int{2, 3, 4}},
		{[]int{1, 2, 3, 4, 5}, -2, 3, []int{4, 5}},
		{[]int{1, 2, 3, 4, 5}, 0, 2, []int{1, 2}},
		{[]int{1, 2, 3, 4, 5}, 2, 5, []int{3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, -6, 2, nil},         // out of range
		{[]int{1, 2, 3, 4, 5}, 3, 10, []int{4, 5}}, // length beyond array length
		{[]int{}, 0, 1, []int{}},                   // empty array
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("start:%d length:%d", tt.start, tt.length), func(t *testing.T) {
			result := utils.Slice(tt.arr, tt.start, tt.length)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestIsset(t *testing.T) {
	tests := []struct {
		array    map[interface{}]interface{}
		key      interface{}
		expected bool
	}{
		{map[interface{}]interface{}{"a": 1, "b": 2}, "a", true},
		{map[interface{}]interface{}{"a": 1, "b": 2}, "c", false},
		{map[interface{}]interface{}{"a": 1, "b": 2}, 1, false},
		{map[interface{}]interface{}{"a": 1, "b": 2}, "b", true},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("key:%v", tt.key), func(t *testing.T) {
			result := utils.Isset(tt.array, tt.key)
			assert.Equal(t, tt.expected, result)
		})
	}
}
