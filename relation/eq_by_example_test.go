package relation_test

import (
	"fmt"
	"math"

	"github.com/xuender/rg/relation"
)

func ExampleEqBy() {
	fmt.Println(relation.EqBy(math.Abs, 5)(-5))

	// Output:
	// true
}
