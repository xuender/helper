package talmud_test

import (
	"testing"

	"github.com/xuender/helper/talmud"
)

func TestDichotomy(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		total float64
		debt1 float64
		debt2 float64
		want  float64
		want2 float64
	}{
		{"case1", 280, 100, 200, 90, 190},
		{"case2", 120, 100, 200, 50, 70},
		{"case3", 20, 100, 200, 10, 10},
		{"case4", 301, 100, 200, 100, 200},
		{"case5", 120, 200, 100, 70, 50},
		{"case6", 200, 100, 500, 50, 150},
		{"case7", 150, 200, 300, 75, 75},
		{"case8", -280, 100, 200, -90, -190},
		{"case9", 0, 100, 200, 0, 0},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			got, got2 := talmud.Dichotomy(testCase.total, testCase.debt1, testCase.debt2)

			if got != testCase.want {
				t.Errorf("Dichotomy() case:%s 1= %v, want %v", testCase.name, got, testCase.want)
			}

			if got2 != testCase.want2 {
				t.Errorf("Dichotomy() case:%s 2= %v, want %v", testCase.name, got2, testCase.want2)
			}
		})
	}
}
