package str

import "regexp"

func Match(regStr string) func(string) []string {
	reg := regexp.MustCompile(regStr)

	return func(str string) []string {
		return reg.FindAllString(str, -1)
	}
}
