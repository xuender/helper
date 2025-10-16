package talmud

import "github.com/xuender/helper/gtype"

type debtWithIndex[N gtype.Number] struct {
	debt  N
	index int
}
