package logic_test

import (
	"fmt"

	"github.com/xuender/ramdago/logic"
)

func ExampleBoth() {
	leapYear := logic.Both(
		func(year int) bool { return year%4 == 0 },
		func(year int) bool { return year%100 != 0 },
	)

	fmt.Println(leapYear(2000))
	fmt.Println(leapYear(2004))

	// Output:
	// false
	// true
}

func ExampleAllPass() {
	leapYear := logic.AllPass(
		func(year int) bool { return year%4 == 0 },
		func(year int) bool { return year%100 != 0 },
	)

	fmt.Println(leapYear(2000))
	fmt.Println(leapYear(2004))

	// Output:
	// false
	// true
}
