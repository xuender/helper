package cont_test

import (
	"fmt"

	"github.com/xuender/helper/cont"
)

func ExampleMin() {
	fmt.Println(cont.Min([]int{}))
	fmt.Println(cont.Min([]int{1}))
	fmt.Println(cont.Min([]int{3, 2}))

	// Output:
	// 0
	// 1
	// 2
}
