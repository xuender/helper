// nolint: nonamedreturns
package ops_test

import (
	"fmt"
	"io"
	"os"

	"github.com/xuender/helper/ops"
)

func callNum() (res error) {
	defer ops.Recover(func(err error) {
		res = err
	})

	panic(3)
}

func callString() (res error) {
	defer ops.Recover(func(err error) {
		res = err
	})

	panic("str")
}

func callError() (res error) {
	defer ops.Recover(func(err error) {
		res = err
	})

	panic(os.ErrExist)
}

func callNoRecover() (res error) {
	defer ops.Recover(func(err error) {
		res = err
	})

	res = io.EOF

	return
}

func ExampleRecover() {
	fmt.Println(callNoRecover())
	fmt.Println(callString())
	fmt.Println(callError())
	fmt.Println(callNum())

	// Output:
	// EOF
	// str
	// file already exists
	// 3
}
