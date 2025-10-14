package lang

import (
	"strconv"
)

// ParseInt converts a string to an int64 value.
// It returns an error if the conversion fails.
func ParseInt(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}

// ParseUint converts a string to a uint64 value.
// It returns an error if the conversion fails.
func ParseUint(str string) (uint64, error) {
	return strconv.ParseUint(str, 10, 64)
}
