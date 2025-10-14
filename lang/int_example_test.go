package lang_test

import (
	"fmt"

	"github.com/xuender/helper/lang"
)

func ExampleParseInt() {
	fmt.Println(lang.ParseInt("-33"))

	// Output:
	// -33 <nil>
}

func ExampleParseUint() {
	fmt.Println(lang.ParseUint("33"))

	// Output:
	// 33 <nil>
}
