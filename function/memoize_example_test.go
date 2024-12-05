package function_test

import (
	"fmt"
	"sync"

	"github.com/xuender/rg/function"
)

func ExampleMemoize() {
	count := 0
	double := func(val int) int {
		count++

		return val * 2
	}
	memoDouble := function.Memoize(double)

	for idx := range 1000 {
		memoDouble(idx % 100)
	}

	fmt.Println(count)

	// Output:
	// 100
}

func ExampleMemoizeSync() {
	count := 0
	double := func(val int) int {
		count++

		return val * 2
	}
	memoDouble := function.MemoizeSync(double)

	group := sync.WaitGroup{}
	group.Add(1000)

	for idx := range 1000 {
		go func() {
			memoDouble(idx % 100)
			group.Done()
		}()
	}

	group.Wait()

	fmt.Println(count)

	// Output:
	// 100
}
