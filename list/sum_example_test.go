package list_test

import (
	"fmt"

	"github.com/xuender/rg/list"
)

func ExampleSum() {
	fmt.Println(list.Sum([]int{1, 2, 3, 4}))

	// Output:
	// 10
}
