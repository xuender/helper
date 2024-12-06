package fp_test

import (
	"fmt"

	"github.com/xuender/helper/fp"
)

func ExamplePipe() {
	work := fp.Pipe(
		func(a int) int { return a + 2 },
		func(a int) int { return a * 3 },
	)

	fmt.Println(work(1))

	// Output:
	// 9
}
