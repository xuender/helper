package relation_test

import (
	"fmt"

	"github.com/xuender/ramdago/relation"
)

func ExampleLte() {
	fmt.Println(relation.Lte(2)(2))
	fmt.Println(relation.Lte('a')('z'))

	// Output:
	// true
	// true
}
