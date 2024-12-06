package slice_test

import (
	"fmt"

	"github.com/xuender/helper/slice"
)

func ExampleContains() {
	fmt.Println(slice.Contains(3)([]int{1, 2, 3}))
	fmt.Println(slice.Contains(4)([]int{1, 2, 3}))
	fmt.Println(slice.Contains([1]int{42})([][1]int{{42}}))

	// Output:
	// true
	// false
	// true
}
