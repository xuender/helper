package function_test

import (
	"fmt"

	"github.com/xuender/rg/function"
	"github.com/xuender/rg/list"
	"github.com/xuender/rg/math"
	"github.com/xuender/rg/relation"
)

func ExampleFlipCheck() {
	lessThan0 := function.FlipCheck(relation.Lt[int])(0)
	lessThan2 := function.FlipCheck(relation.Lt[int])(2)

	fmt.Println(list.Any(lessThan0)([]int{1, 2}))
	fmt.Println(list.Any(lessThan2)([]int{1, 2}))

	// Output:
	// false
	// true
}

func ExampleFlipMap() {
	gToKg := function.FlipMap(math.Divide[float64])(1_000)

	fmt.Printf("%.2fKG", gToKg(382))

	// Output:
	// 0.38KG
}
