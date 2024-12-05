package relation_test

import (
	"fmt"

	"github.com/xuender/rg/relation"
)

func ExampleGte() {
	fmt.Println(relation.Gte(2)(2))
	fmt.Println(relation.Gte('a')('z'))

	// Output:
	// true
	// false
}
