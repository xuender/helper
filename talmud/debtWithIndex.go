package talmud

import "github.com/xuender/helper/types"

type debtWithIndex[N types.Number] struct {
	debt  N
	index int
}
