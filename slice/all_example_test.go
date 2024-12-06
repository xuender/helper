package slice_test

import (
	"fmt"

	"github.com/xuender/helper/slice"
)

func ExampleAll() {
	equals3 := func(num int) bool { return num == 3 }

	fmt.Println(slice.All([]int{3, 3, 3, 3}, equals3))
	fmt.Println(slice.All([]int{3, 3, 1, 3}, equals3))

	// Output:
	// true
	// false
}
