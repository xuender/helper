package cont_test

import (
	"fmt"

	"github.com/xuender/helper/cont"
)

func ExampleDifference() {
	fmt.Println(cont.Difference([]int{1, 2}))
	fmt.Println(cont.Difference([]int{1, 2}, []int{2, 3}))

	// Output:
	// [1 2]
	// [1]
}

func ExampleDifferenceFunc() {
	equal := func(a, b int) bool { return a == b }
	fmt.Println(cont.DifferenceFunc(equal, []int{1, 2}))
	fmt.Println(cont.DifferenceFunc(equal, []int{1, 2}, []int{2, 3}))

	// Output:
	// [1 2]
	// [1]
}
