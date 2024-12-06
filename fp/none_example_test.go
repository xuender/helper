package fp_test

import (
	"fmt"

	"github.com/xuender/helper/fp"
)

func ExampleNone() {
	equals3 := func(num int) bool { return num == 3 }

	fmt.Println(fp.None(equals3)([]int{1, 2, 4, 5}))
	fmt.Println(fp.None(equals3)([]int{3, 3, 1, 3}))

	// Output:
	// true
	// false
}
