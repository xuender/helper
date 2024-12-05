package function_test

import (
	"fmt"

	"github.com/xuender/rg/function"
)

func ExampleComplement() {
	gt10 := func(val int) bool { return val > 10 }
	lte10 := function.Complement(gt10)

	fmt.Println(lte10(5))

	// Output:
	// true
}
