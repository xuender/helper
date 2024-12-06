package fp_test

import (
	"fmt"

	"github.com/xuender/helper/fp"
)

func ExampleAny() {
	greaterThan0 := func(num int) bool { return 0 > num }
	greaterThan2 := func(num int) bool { return 2 > num }

	fmt.Println(fp.Any(greaterThan0)([]int{1, 2}))
	fmt.Println(fp.Any(greaterThan2)([]int{1, 2}))

	// Output:
	// false
	// true
}
