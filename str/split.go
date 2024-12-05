package str

import "strings"

// Split returns a function that splits a string by a specified separator.
func Split(sep string) func(string) []string {
	return func(data string) []string {
		return strings.Split(data, sep)
	}
}
