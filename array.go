package utils

import "strings"

func InArray[T comparable](needle T, haystack []T) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}
	return false
}

func Explode(separator string, stringStr string) []string {
	return strings.Split(stringStr, separator)
}

func Implode(separator string, elements []string) string {
	return strings.Join(elements, separator)
}
