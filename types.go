package utils

import (
	"math"
	"strconv"
)

func DetectType(input string) interface{} {
	if i, err := strconv.ParseInt(input, 10, 64); err == nil {
		if i >= math.MinInt32 && i <= math.MaxInt32 {
			return int32(i)
		}
		return int64(i)
	}

	if f, err := strconv.ParseFloat(input, 64); err == nil {
		if f >= -math.MaxFloat32 && f <= math.MaxFloat32 {
			return float32(f)
		}
		return f
	}

	if b, err := strconv.ParseBool(input); err == nil {
		return b
	}

	return input
}

func IsInt(input string) (bool, int) {
	i, err := strconv.Atoi(input)
	if err != nil || i < math.MinInt32 || i > math.MaxInt32 {
		return false, 0
	}
	return true, i
}
