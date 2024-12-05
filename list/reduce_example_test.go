package list_test

import (
	"fmt"

	"github.com/xuender/rg/list"
	"github.com/xuender/rg/math"
)

func ExampleReduce() {
	adder := list.Reduce(func(val1, val2 int) int {
		return val1 + val2
	})
	fmt.Println(adder([]int{1, 2, 3, 4}))

	// Output:
	// 10
}

func ExampleReduceCurry() {
	adder := list.ReduceCurry(math.Add[int])
	fmt.Println(adder([]int{1, 2, 3, 4}))

	// Output:
	// 10
}
