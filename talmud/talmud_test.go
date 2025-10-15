package talmud_test

import (
	"fmt"

	"github.com/xuender/helper/talmud"
)

func ExampleTalmud() {
	fmt.Println(talmud.Talmud(150))
	fmt.Println(talmud.Talmud(0, 100, 200, 300))
	fmt.Println(talmud.Talmud(150, 100, 200, 300))
	fmt.Println(talmud.Talmud(200, 100, 200, 300))
	fmt.Println(talmud.Talmud(300, 100, 200, 300))
	fmt.Println(talmud.Talmud(600, 100, 200, 300))
	fmt.Println(talmud.Talmud(1200, 300, 200, 100))
	// fmt.Println(talmud.Talmud(-1200, 300, 200, 100))
	fmt.Println(talmud.Talmud(300, 100, 120, 200, 300))
	// fmt.Println(talmud.Talmud(-300, 100, 200, 300))
	fmt.Println(talmud.Talmud(0, 100, 200, 300))

	// OutPut:
	// []
	// [0 0 0]
	// [50 50 50]
	// [50 75 75]
	// [50 100 150]
	// [100 200 300]
	// [300 200 100]
	// [50 60 95 95]
	// [0 0 0]
}

func ExampleTalmud_out() {
	fmt.Println(talmud.Talmud(280, 100, 200))
	fmt.Println(talmud.Talmud(350, 100, 120, 150, 200))
	fmt.Println(talmud.Talmud(40, 100, 120, 150, 200))
	fmt.Println(talmud.Talmud(280, 100))
	fmt.Println(talmud.Talmud(280, 300))

	// OutPut:
	// [90 190]
	// [50 60 95 145]
	// [10 10 10 10]
	// [100]
	// [280]
}
