package relation_test

import (
	"fmt"

	"github.com/xuender/rg/relation"
)

func ExampleEquals() {
	fmt.Println(relation.Equals(1)(1))
	fmt.Println(relation.Equals(1)('1'))
	fmt.Println(relation.Equals([3]int{1, 2, 3})([3]int{1, 2, 3}))

	// Output:
	// true
	// false
	// true
}
