package helper_test

import (
	"fmt"

	"github.com/xuender/helper"
)

func ExampleEither() {
	check := helper.Either(
		func(val int) bool { return val > 10 },
		func(val int) bool { return val%2 == 0 },
	)

	fmt.Println(check(101))
	fmt.Println(check(8))
	fmt.Println(check(7))

	// Output:
	// true
	// true
	// false
}
