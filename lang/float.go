package lang

import (
	"strconv"
)

// ParseFloat converts a string to a float64 value.
// It returns an error if the conversion fails.
func ParseFloat(str string) (float64, error) {
	return strconv.ParseFloat(str, 64)
}
