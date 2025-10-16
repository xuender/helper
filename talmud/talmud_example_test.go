package talmud_test

import (
	"fmt"

	"github.com/xuender/helper/talmud"
)

func ExampleTalmud() {
	fmt.Println(talmud.Talmud(150, []int{100, 200, 300}))
	fmt.Println(talmud.Talmud(200, []int{100, 200, 300}))
	fmt.Println(talmud.Talmud(300, []int{100, 200, 300}))
	fmt.Println(talmud.Talmud(600, []int{100, 200, 300}))
	fmt.Println(talmud.Talmud(1200, []int{300, 200, 100}))
	fmt.Println(talmud.Talmud(-1200, []int{300, 200, 100}))

	// Output:
	// [50 50 50]
	// [50 75 75]
	// [50 100 150]
	// [100 200 300]
	// [300 200 100]
	// [-300 -200 -100]
}
