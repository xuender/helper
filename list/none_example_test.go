package list_test

import (
	"fmt"

	"github.com/xuender/rg/list"
	"github.com/xuender/rg/relation"
)

func ExampleNone() {
	equals3 := relation.Equals(3)
	fmt.Println(list.None(equals3)([]int{1, 2, 4, 5}))
	fmt.Println(list.None(equals3)([]int{3, 3, 1, 3}))

	// Output:
	// true
	// false
}
