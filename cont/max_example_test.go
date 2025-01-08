package cont_test

import (
	"fmt"

	"github.com/xuender/helper/cont"
)

func ExampleMax() {
	fmt.Println(cont.Max[int](nil))
	fmt.Println(cont.Max([]int{}))
	fmt.Println(cont.Max([]int{1}))
	fmt.Println(cont.Max([]int{2, 3}))

	// Output:
	// 0
	// 0
	// 1
	// 3
}
