package slice_test

import (
	"fmt"

	"github.com/xuender/helper/slice"
)

func ExampleAny() {
	greaterThan0 := func(num int) bool { return 0 > num }
	greaterThan2 := func(num int) bool { return 2 > num }

	fmt.Println(slice.Any(greaterThan0)([]int{1, 2}))
	fmt.Println(slice.Any(greaterThan2)([]int{1, 2}))

	// Output:
	// false
	// true
}
