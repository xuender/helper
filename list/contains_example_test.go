package list_test

import (
	"fmt"

	"github.com/xuender/rg/list"
)

func ExampleContains() {
	fmt.Println(list.Contains(3)([]int{1, 2, 3}))
	fmt.Println(list.Contains(4)([]int{1, 2, 3}))
	fmt.Println(list.Contains([1]int{42})([][1]int{{42}}))

	// Output:
	// true
	// false
	// true
}
