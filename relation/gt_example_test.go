package relation_test

import (
	"fmt"

	"github.com/xuender/rg/relation"
)

func ExampleGt() {
	fmt.Println(relation.Gt(2)(1))
	fmt.Println(relation.Gt('a')('z'))

	// Output:
	// true
	// false
}
