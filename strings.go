package utils

import "strings"

func Substr(s string, start int, length *int) string {
	if start < 0 {
		start = len(s) + start
	}
	if start < 0 || start >= len(s) {
		return ""
	}

	var end int
	if length == nil {
		end = len(s)
	} else {
		end = start + *length
	}

	if end > len(s) {
		end = len(s)
	}
	return s[start:end]
}

func Strpos(stringStr string, needle string) int {
	return strings.LastIndex(stringStr, needle)
}
