package slice_test

import (
	"fmt"

	"github.com/xuender/helper/slice"
)

func ExampleMin() {
	fmt.Println(slice.Min[int](nil))
	fmt.Println(slice.Min([]int{}))
	fmt.Println(slice.Min([]int{1}))
	fmt.Println(slice.Min([]int{3, 2}))

	// Output:
	// 0
	// 0
	// 1
	// 2
}
