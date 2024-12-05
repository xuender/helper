package list_test

import (
	"fmt"

	"github.com/xuender/rg/list"
)

func ExampleZipWith() {
	adder := func(num1, num2 int) int { return num1 + num2 }
	zip := list.ZipWith(adder, []int{1, 3, 5, 7, 9})

	fmt.Println(zip([]int{2, 4, 6, 8}))

	// Output:
	// [3 7 11 15]
}
