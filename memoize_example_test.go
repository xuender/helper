package helper_test

import (
	"fmt"

	"github.com/xuender/helper"
)

func ExampleMemoize() {
	count := 0
	double := func(val int) int {
		count++

		return val * 2
	}
	memoDouble := helper.Memoize(double)

	for idx := range 1000 {
		memoDouble(idx % 100)
	}

	fmt.Println(count)

	// Output:
	// 100
}
