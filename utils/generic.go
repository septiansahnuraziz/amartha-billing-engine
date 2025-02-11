package utils

import (
	"strconv"
	"strings"
)

type Int interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 | uintptr |
		float32 | float64
}

// ValueOrDefault returns the provided value if it is not empty, otherwise returns the defaultValue.
func ValueOrDefault[T comparable](value, defaultValue T) T {
	var emptyValue T

	if value == emptyValue {
		return defaultValue
	}

	return value
}

// StringToInt converts a string to an integer value.
func StringToInt[T Int](s string) T {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return T(i)
}

func MapToStruct[T any](value any) (item T, err error) {
	data, err := JSONMarshal(value)
	if err != nil {
		return item, err
	}

	if err := JSONUnmarshal(data, &item); err != nil {
		return item, err
	}

	return item, nil
}
func StringToIntOrFloat[T Int](s string) T {
	// Replace commas with dots for float parsing
	s = strings.ReplaceAll(s, ",", ".")

	// Try to parse as float
	f, err := strconv.ParseFloat(s, 64)
	if err == nil {
		return T(f)
	}

	// Try to parse as integer if float parsing fails
	i, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		return T(i)
	}

	return 0
}

func ExpectedNumber[T Int](v any) T {
	var result T
	switch value := v.(type) {
	case int:
		result = T(value)
	case int8:
		result = T(value)
	case int16:
		result = T(value)
	case int32:
		result = T(value)
	case int64:
		result = T(value)
	case uint:
		result = T(value)
	case uint8:
		result = T(value)
	case uint16:
		result = T(value)
	case uint32:
		result = T(value)
	case uint64:
		result = T(value)
	case uintptr:
		result = T(value)
	case float32:
		result = T(value)
	case float64:
		result = T(value)
	case string:
		result = T(StringToInt[T](value))
	default:
		result = 0
	}

	return T(result)
}
