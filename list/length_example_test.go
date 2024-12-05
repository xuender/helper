package list_test

import (
	"fmt"

	"github.com/xuender/rg/list"
)

func ExampleLength() {
	fmt.Println(list.Length([]int{1, 2, 3}))

	// Output:
	// 3
}
