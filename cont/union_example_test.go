package cont_test

import (
	"fmt"

	"github.com/xuender/helper/cont"
)

func ExampleUnion() {
	fmt.Println(cont.Union[[]int]())
	fmt.Println(cont.Union([]int{1, 1}))
	fmt.Println(cont.Union([]int{1, 2}, []int{2, 3}))

	// Output:
	// []
	// [1]
	// [1 2 3]
}

func ExampleUnionFunc() {
	equal := func(a, b int) bool { return a == b }

	fmt.Println(cont.UnionFunc[[]int](equal))
	fmt.Println(cont.UnionFunc(equal, []int{1, 1}))
	fmt.Println(cont.UnionFunc(equal, []int{1, 2}, []int{2, 3}))

	// Output:
	// []
	// [1]
	// [1 2 3]
}
