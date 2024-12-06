package helper_test

import (
	"fmt"

	"github.com/xuender/helper"
)

func ExampleMatch() {
	fmt.Println(helper.Match("[a-z]a")("bananas"))
	fmt.Println(helper.Match("a")("b"))

	// Output:
	// [ba na na]
	// []
}
