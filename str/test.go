package str

import "regexp"

func Test(regStr string) func(string) bool {
	reg := regexp.MustCompile(regStr)

	return func(str string) bool {
		return reg.MatchString(str)
	}
}
