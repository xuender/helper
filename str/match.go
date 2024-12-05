package str

import "regexp"

// Match returns a function that finds all matches of a regex pattern in a string.
func Match(regStr string) func(string) []string {
	reg := regexp.MustCompile(regStr)

	return func(str string) []string {
		return reg.FindAllString(str, -1)
	}
}
