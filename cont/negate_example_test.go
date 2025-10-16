package cont_test

import (
	"fmt"

	"github.com/xuender/helper/cont"
)

func ExampleNegate() {
	fmt.Println(cont.Negate([]int{1, 2, 3}))
	fmt.Println(cont.Negate([]int{1, -2, 3}))

	// Output:
	// [-1 -2 -3]
	// [-1 2 -3]
}
