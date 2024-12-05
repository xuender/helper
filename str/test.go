package str

import "regexp"

// Test returns a function that checks if a string matches a regex pattern.
func Test(regStr string) func(string) bool {
	reg := regexp.MustCompile(regStr)

	return func(str string) bool {
		return reg.MatchString(str)
	}
}
