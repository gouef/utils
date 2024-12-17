package utils

import (
	"strings"
)

func InArray[T comparable](needle T, haystack []T) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}
	return false
}

func Explode(separator string, stringStr string) []string {
	if separator == "" {
		return []string{stringStr}
	}

	return strings.Split(stringStr, separator)
}

func Implode(separator string, elements []string) string {
	return strings.Join(elements, separator)
}

func Slice(arr []int, start, length int) []int {
	if start < 0 {
		start = len(arr) + start
	}
	end := start + length
	if start < 0 || start > len(arr) || end < 0 {
		return nil
	}
	if end > len(arr) {
		end = len(arr)
	}
	return arr[start:end]
}

func Isset(array map[interface{}]interface{}, key any) bool {
	_, ok := array[key]
	return ok
}
