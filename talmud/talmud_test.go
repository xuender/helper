package talmud_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/helper/talmud"
)

func TestTalmud(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)

	ass.Nil(talmud.Talmud(150, nil))

	cases := []struct {
		total  int
		debts  []int
		repays []int
	}{
		{0, []int{100, 200, 300}, []int{0, 0, 0}},
		{50, []int{100, 200, 300}, []int{16, 16, 16}},
		{100, []int{100, 200, 300}, []int{33, 33, 33}},
		{150, []int{100, 200, 300}, []int{50, 50, 50}},
		{200, []int{100, 200, 300}, []int{50, 75, 75}},
		{250, []int{100, 200, 300}, []int{50, 100, 100}},
		{300, []int{100, 200, 300}, []int{50, 100, 150}},
		{350, []int{100, 200, 300}, []int{50, 100, 200}},
		{400, []int{100, 200, 300}, []int{50, 125, 225}},
		{450, []int{100, 200, 300}, []int{50, 150, 250}},
		{500, []int{100, 200, 300}, []int{50, 175, 275}},
		{550, []int{100, 200, 300}, []int{75, 187, 287}},
		{600, []int{100, 200, 300}, []int{100, 200, 300}},
		{650, []int{100, 200, 300}, []int{100, 200, 300}},
		{1200, []int{100, 200, 300}, []int{100, 200, 300}},
		{-1200, []int{300, 200, 100}, []int{-300, -200, -100}},
		{300, []int{100, 120, 200, 300}, []int{50, 60, 95, 95}},
		{-300, []int{100, 200, 300}, []int{-50, -100, -150}},
		{280, []int{100, 200}, []int{90, 190}},
		{350, []int{100, 120, 150, 200}, []int{50, 60, 95, 145}},
		{40, []int{100, 120, 150, 200}, []int{10, 10, 10, 10}},
		{280, []int{100}, []int{100}},
		{280, []int{300}, []int{280}},
	}

	for index, testCase := range cases {
		name := fmt.Sprintf("case %d", index+1)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			repays := talmud.Talmud(testCase.total, testCase.debts)
			ass.Equal(testCase.repays, repays, name)
		})
	}
}
