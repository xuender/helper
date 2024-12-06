package fp_test

import (
	"fmt"

	"github.com/xuender/helper/fp"
)

func ExampleAll() {
	equals3 := func(num int) bool { return num == 3 }

	fmt.Println(fp.All(equals3)([]int{3, 3, 3, 3}))
	fmt.Println(fp.All(equals3)([]int{3, 3, 1, 3}))

	// Output:
	// true
	// false
}
