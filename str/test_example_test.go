package str_test

import (
	"fmt"

	"github.com/xuender/rg/str"
)

func ExampleTest() {
	fmt.Println(str.Test("^x")("xyz"))
	fmt.Println(str.Test("^y")("xyz"))

	// Output:
	// true
	// false
}
