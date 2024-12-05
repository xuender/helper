package logic_test

import (
	"fmt"

	"github.com/xuender/ramdago/logic"
)

func ExampleEither() {
	check := logic.Either(
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
