package fp_test

import (
	"fmt"

	"github.com/xuender/helper/fp"
)

func ExampleReduce() {
	adder := fp.Reduce(func(val1, val2 int) int {
		return val1 + val2
	})
	fmt.Println(adder([]int{1, 2, 3, 4}))

	// Output:
	// 10
}
