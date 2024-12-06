package fp_test

import (
	"fmt"

	"github.com/xuender/helper/fp"
)

func ExampleBoth() {
	leapYear := fp.Both(
		func(year int) bool { return year%4 == 0 },
		func(year int) bool { return year%100 != 0 },
	)

	fmt.Println(leapYear(2000))
	fmt.Println(leapYear(2004))

	// Output:
	// false
	// true
}
