package list_test

import (
	"fmt"

	"github.com/xuender/rg/list"
)

func ExampleMin() {
	fmt.Println(list.Min([]int{2, 1, 3}))
	fmt.Println(list.Min([]int{}))

	// Output:
	// 1
	// 0
}
