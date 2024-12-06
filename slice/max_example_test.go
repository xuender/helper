package slice_test

import (
	"fmt"

	"github.com/xuender/helper/slice"
)

func ExampleMax() {
	fmt.Println(slice.Max[int](nil))
	fmt.Println(slice.Max([]int{}))
	fmt.Println(slice.Max([]int{1}))
	fmt.Println(slice.Max([]int{2, 3}))

	// Output:
	// 0
	// 0
	// 1
	// 3
}
