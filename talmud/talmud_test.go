package talmud_test

import (
	"fmt"

	"github.com/xuender/helper/talmud"
)

func ExampleTalmud() {
	fmt.Println(talmud.Talmud(150, nil))
	fmt.Println(talmud.Talmud(0, []int{100, 200, 300}))
	fmt.Println(talmud.Talmud(150, []int{100, 200, 300}))
	fmt.Println(talmud.Talmud(200, []int{100, 200, 300}))
	fmt.Println(talmud.Talmud(300, []int{100, 200, 300}))
	fmt.Println(talmud.Talmud(600, []int{100, 200, 300}))
	fmt.Println(talmud.Talmud(1200, []int{300, 200, 100}))
	fmt.Println(talmud.Talmud(-1200, []int{300, 200, 100}))
	fmt.Println(talmud.Talmud(300, []int{100, 120, 200, 300}))
	fmt.Println(talmud.Talmud(-300, []int{100, 200, 300}))
	fmt.Println(talmud.Talmud(280, []int{100, 200}))
	fmt.Println(talmud.Talmud(350, []int{100, 120, 150, 200}))
	fmt.Println(talmud.Talmud(40, []int{100, 120, 150, 200}))
	fmt.Println(talmud.Talmud(280, []int{100}))
	fmt.Println(talmud.Talmud(280, []int{300}))
	fmt.Println(talmud.Talmud(-300, []int{100, 200, 300}))

	// OutPut:
	// []
	// [0 0 0]
	// [50 50 50]
	// [50 75 75]
	// [50 100 150]
	// [100 200 300]
	// [300 200 100]
	// [-300 -200 -100]
	// [50 60 95 95]
	// [-50 -100 -150]
	// [90 190]
	// [50 60 95 145]
	// [10 10 10 10]
	// [100]
	// [280]
	// [-50 -100 -150]
}

func ExampleTalmud_for() {
	for num := range 14 {
		total := num * 50
		fmt.Println(total, talmud.Talmud(total, []int{100, 200, 300}))
	}

	// Output:
	// 0 [0 0 0]
	// 50 [16 16 16]
	// 100 [33 33 33]
	// 150 [50 50 50]
	// 200 [50 75 75]
	// 250 [50 100 100]
	// 300 [50 100 150]
	// 350 [50 100 200]
	// 400 [50 125 225]
	// 450 [50 150 250]
	// 500 [50 175 275]
	// 550 [75 187 287]
	// 600 [100 200 300]
	// 650 [100 200 300]
}
