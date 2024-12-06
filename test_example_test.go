package helper_test

import (
	"fmt"

	"github.com/xuender/helper"
)

func ExampleTest() {
	fmt.Println(helper.Test("^x")("xyz"))
	fmt.Println(helper.Test("^y")("xyz"))

	// Output:
	// true
	// false
}
