package function_test

import (
	"fmt"

	"github.com/xuender/rg/function"
)

func ExampleTap() {
	tapNum := function.Tap(func(num int) { fmt.Println(num * 2) })

	fmt.Println(tapNum(3))

	// Output:
	// 6
	// 3
}
