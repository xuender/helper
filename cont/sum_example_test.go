package cont_test

import (
	"fmt"

	"github.com/xuender/helper/cont"
)

func ExampleSum() {
	fmt.Println(cont.Sum([]int{1, 2, 3, 4}))

	// Output:
	// 10
}
