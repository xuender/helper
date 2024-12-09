package slice_test

import (
	"fmt"

	"github.com/xuender/helper/slice"
)

func ExampleFirst() {
	num, has := slice.First([]int{1, 2, 3}, func(item int) bool {
		return item%2 == 0
	})

	fmt.Println(num, has)

	num, has = slice.First([]int{1, 2, 3}, func(item int) bool {
		return item%7 == 0
	})

	fmt.Println(num, has)

	// Output:
	// 2 true
	// 0 false
}
