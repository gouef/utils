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

func Substring(s string, start int, length *int) string {
	return Substr(s, start, length)
}

func Strpos(stringStr string, needle string) int {
	return strings.LastIndex(stringStr, needle)
}

func StringPosition(stringStr string, needle string) int {
	return Strpos(stringStr, needle)
}

func Strncmp(s1, s2 string, n int) int {
	sub1 := s1
	sub2 := s2

	if len(s1) > n {
		sub1 = s1[:n]
	}
	if len(s2) > n {
		sub2 = s2[:n]
	}

	return strings.Compare(sub1, sub2)
}

func StringCompare(s1, s2 string, n int) int {
	return Strncmp(s1, s2, n)
}

func StringJoin(values []string) string {
	return strings.Join(values, " ")
}
