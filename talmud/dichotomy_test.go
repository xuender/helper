package talmud_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/helper/talmud"
)

func TestDichotomy(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	cases := []struct {
		total  int
		debt1  int
		debt2  int
		repay1 int
		repay2 int
	}{
		{280, 100, 200, 90, 190},
		{120, 100, 200, 50, 70},
		{20, 100, 200, 10, 10},
		{301, 100, 200, 100, 200},
		{120, 200, 100, 70, 50},
		{200, 100, 500, 50, 150},
		{150, 200, 300, 75, 75},
		{-280, 100, 200, -90, -190},
		{0, 100, 200, 0, 0},
	}

	for index, testCase := range cases {
		name := fmt.Sprintf("case: %d", index+1)

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			repay1, repay2 := talmud.Dichotomy(testCase.total, testCase.debt1, testCase.debt2)

			ass.Equal(testCase.repay1, repay1, name)
			ass.Equal(testCase.repay2, repay2, name)
		})
	}
}
