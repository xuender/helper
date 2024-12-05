package function_test

import (
	"fmt"

	"github.com/xuender/rg/function"
	"github.com/xuender/rg/list"
	"github.com/xuender/rg/math"
	"github.com/xuender/rg/types"
)

func ExampleConverge() {
	average := function.Converge(
		func(sum, length int) int { return sum / length },
		[]types.ReducerSlice[int]{list.Sum[int], list.Length[int]},
	)

	fmt.Println(average([]int{1, 2, 3, 4, 5, 6, 7}))

	// Output:
	// 4
}

func ExampleConvergeCurry() {
	average := function.ConvergeCurry(
		math.Divide[int],
		[]types.ReducerSlice[int]{list.Sum[int], list.Length[int]},
	)

	fmt.Println(average([]int{1, 2, 3, 4, 5, 6, 7}))

	// Output:
	// 4
}
