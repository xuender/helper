package slice_test

import (
	"fmt"

	"github.com/xuender/helper/slice"
)

func ExampleSum() {
	fmt.Println(slice.Sum([]int{1, 2, 3, 4}))

	// Output:
	// 10
}
