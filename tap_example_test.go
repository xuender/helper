package helper_test

import (
	"fmt"

	"github.com/xuender/helper"
)

func ExampleTap() {
	tapNum := helper.Tap(func(num int) { fmt.Println(num * 2) })

	fmt.Println(tapNum(3))

	// Output:
	// 6
	// 3
}
