package talmud

import (
	"sort"

	"github.com/xuender/helper/cont"
	"github.com/xuender/helper/gtype"
)

// Talmud implements the Talmudic algorithm for debt repayment distribution.
// It distributes a total amount among multiple debts according to Talmudic principles.
//
// Parameters:
//   - total: the total amount available for distribution
//   - debts: a slice of debt amounts to be repaid
//
// Returns:
//   - A slice of numbers representing the actual repayment amounts for each debt,
//     corresponding to the order of the input debts
//
// nolint: nonamedreturns
func Talmud[N gtype.Number](total N, debts []N) (repays []N) {
	length := len(debts)
	if total == 0 {
		return make([]N, length)
	}

	switch length {
	case 0:
		return nil
	case 1:
		return []N{cont.Min([]N{debts[0], total})}
	case _two:
		repay1, repay2 := Dichotomy(total, debts[0], debts[1])

		return []N{repay1, repay2}
	}

	repays = make([]N, length)

	if total < 0 {
		defer func() {
			repays = cont.Negate(repays)
		}()

		total = -total
	}

	sortedDebts := make([]debtWithIndex[N], length)
	for index, debt := range debts {
		sortedDebts[index] = debtWithIndex[N]{debt: debt, index: index}
	}

	sort.Slice(sortedDebts, func(i, j int) bool {
		return sortedDebts[i].debt < sortedDebts[j].debt
	})

	if total <= sortedDebts[0].debt*N(length)/_two {
		repayEach(total, length, repays)

		return repays
	}

	for index, current := range sortedDebts[:length-1] {
		var totalOther N

		for _, debt := range sortedDebts[index+1:] {
			totalOther += debt.debt
		}

		repays[current.index], total = Dichotomy(total, current.debt, totalOther)
	}

	repays[sortedDebts[length-1].index] = total

	return repays
}

func repayEach[N gtype.Number](total N, length int, repays []N) {
	each := total / N(length)

	for i := range length {
		repays[i] = each
	}
}
