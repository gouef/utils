package utils

import "strconv"

func DetectType(input string) interface{} {
	if i, err := strconv.Atoi(input); err == nil {
		return i
	}

	if i, err := strconv.ParseInt(input, 10, 32); err == nil {
		return int32(i)
	}

	if i, err := strconv.ParseInt(input, 10, 64); err == nil {
		return int64(i)
	}

	if f, err := strconv.ParseFloat(input, 64); err == nil {
		return f
	}

	if f, err := strconv.ParseFloat(input, 32); err == nil {
		return float32(f)
	}

	if b, err := strconv.ParseBool(input); err == nil {
		return b
	}

	return input
}

func IsInt(input string) (bool, int) {
	i, err := strconv.Atoi(input)
	return err == nil, i
}
