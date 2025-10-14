package talmud_test

import (
	"fmt"

	"github.com/xuender/helper/talmud"
)

func ExampleTalmud() {
	fmt.Println(talmud.Talmud(150, 100, 200, 300))
	fmt.Println(talmud.Talmud(200, 100, 200, 300))
	fmt.Println(talmud.Talmud(300, 100, 200, 300))
	fmt.Println(talmud.Talmud(600, 100, 200, 300))
	fmt.Println(talmud.Talmud(1200, 300, 200, 100))
	fmt.Println(talmud.Talmud(-1200, 300, 200, 100))
	fmt.Println(talmud.Talmud(300, 100, 120, 200, 300))
	fmt.Println(talmud.Talmud(-300, 100, 200, 300))
	fmt.Println(talmud.Talmud(0, 100, 200, 300))
	fmt.Println(talmud.Talmud(150))

	// OutPut:
	// [50 50 50] 0
	// [50 75 75] 0
	// [50 100 150] 0
	// [100 200 300] 0
	// [300 200 100] 600
	// [-300 -200 -100] -600
	// [50 60 95 95] 0
	// [-50 -100 -150] 0
	// [0 0 0] 0
	// [] 150
}
