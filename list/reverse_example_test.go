package list_test

import (
	"fmt"

	"github.com/xuender/rg/list"
)

func ExampleReverse() {
	fmt.Println(list.Reverse([]int{1, 2, 3}))

	// Output:
	// [3 2 1]
}
