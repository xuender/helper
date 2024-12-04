package str

import "strings"

func Split(sep string) func(string) []string {
	return func(data string) []string {
		return strings.Split(data, sep)
	}
}
