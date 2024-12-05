package list_test

import (
	"fmt"

	"github.com/xuender/rg/list"
	"github.com/xuender/rg/relation"
)

func ExampleAny() {
	greaterThan0 := relation.Gt(0)
	greaterThan2 := relation.Gt(2)

	fmt.Println(list.Any(greaterThan0)([]int{1, 2}))
	fmt.Println(list.Any(greaterThan2)([]int{1, 2}))

	// Output:
	// false
	// true
}
