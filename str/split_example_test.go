package str_test

import (
	"fmt"

	"github.com/xuender/ramdago/str"
)

func ExampleSplit() {
	fmt.Println(str.Split(".")("a.b.c.xyz.d"))

	// Output:
	// [a b c xyz d]
}
