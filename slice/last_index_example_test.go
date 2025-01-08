package slice_test

import (
	"fmt"

	"github.com/xuender/helper/slice"
)

func ExampleLastIndex() {
	fmt.Println(slice.LastIndex([]int{1, 2, 3}, 3))
	fmt.Println(slice.LastIndex([]int{1, 2, 3}, 2))
	fmt.Println(slice.LastIndex([]int{1, 2, 3}, 1))
	fmt.Println(slice.LastIndex([]int{1, 2, 3}, 4))

	// Output:
	// 2
	// 1
	// 0
	// -1
}

func ExampleLastIndexFunc() {
	fmt.Println(slice.LastIndexFunc([]int{1, 2, 3}, func(num int) bool { return num == 2 }))
	fmt.Println(slice.LastIndexFunc([]int{1, 2, 3}, func(num int) bool { return num > 10 }))

	// Output:
	// 1
	// -1
}
