package cont_test

import (
	"fmt"

	"github.com/xuender/helper/cont"
)

func ExampleIntersection() {
	fmt.Println(cont.Intersection[[]int]())
	fmt.Println(cont.Intersection([]int{1, 2, 3}))
	fmt.Println(cont.Intersection([]int{1, 2, 2, 3}, []int{2, 3, 4}))
	fmt.Println(cont.Intersection([]int{1, 1, 2, 3}, []int{2, 3, 4}))

	// Output:
	// []
	// [1 2 3]
	// [2 3]
	// [2 3]
}
