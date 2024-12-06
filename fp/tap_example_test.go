package fp_test

import (
	"fmt"

	"github.com/xuender/helper/fp"
)

func ExampleTap() {
	tapNum := fp.Tap(func(num int) { fmt.Println(num * 2) })

	fmt.Println(tapNum(3))

	// Output:
	// 6
	// 3
}
