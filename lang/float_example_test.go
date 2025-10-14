package lang_test

import (
	"fmt"

	"github.com/xuender/helper/lang"
)

func ExampleParseFloat() {
	fmt.Println(lang.ParseFloat("3.14"))

	// Output:
	// 3.14 <nil>
}
