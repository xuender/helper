package seq_test

import (
	"fmt"
	"time"

	"github.com/xuender/helper/seq"
)

func ExampleCache() {
	input := func(yield func(int) bool) {
		for idx := range 5 {
			fmt.Println("input", idx)
			time.Sleep(time.Millisecond * 100)

			if !yield(idx) {
				return
			}
		}
	}

	for num := range seq.Cache(input, 3) {
		time.Sleep(time.Millisecond * 300)
		fmt.Println("output", num)
	}

	// Output:
	// input 0
	// input 1
	// input 2
	// input 3
	// output 0
	// input 4
	// output 1
	// output 2
	// output 3
	// output 4
}
