package cont_test

import (
	"fmt"

	"github.com/xuender/helper/cont"
)

func ExampleFilter() {
	items := cont.Filter([]int{1, 2, 3, 4}, func(num int) bool { return num%2 == 0 })

	fmt.Println(items)
	fmt.Println(len(items), cap(items))

	// Output:
	// [2 4]
	// 2 2
}
