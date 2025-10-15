package talmud

import (
	"sort"

	"github.com/xuender/helper/cont"
	"github.com/xuender/helper/types"
)

// Talmud implements the Talmudic algorithm for debt repayment distribution.
// It distributes a total amount among multiple debts according to Talmudic principles.
//
// Parameters:
//   - total: the total amount available for distribution
//   - debts: a variadic list of debt amounts to be repaid
//
// Returns:
//   - A slice of numbers representing the actual repayment amounts for each debt,
//     corresponding to the order of the input debts
func Talmud[N types.Number](total N, debts ...N) []N {
	if total == 0 {
		return make([]N, len(debts))
	}

	length := len(debts)

	switch length {
	case 0:
		return nil
	case 1:
		return []N{cont.Min([]N{debts[0], total})}
	case _two:
		repay1, repay2 := Dichotomy(total, debts[0], debts[1])

		return []N{repay1, repay2}
	}

	sortedDebts := make([]debtWithIndex[N], len(debts))
	for index, debt := range debts {
		sortedDebts[index] = debtWithIndex[N]{debt: debt, index: index}
	}

	sort.Slice(sortedDebts, func(i, j int) bool {
		return sortedDebts[i].debt < sortedDebts[j].debt
	})

	repays := make([]N, length)

	if total <= sortedDebts[0].debt/2*N(length) {
		each := total / N(length)

		for i := range length {
			repays[i] = each
		}

		return repays
	}

	for index := range length - 1 {
		var other N

		for _, debt := range sortedDebts[index+1:] {
			other += debt.debt
		}

		repays[sortedDebts[index].index], total = Dichotomy(total, sortedDebts[index].debt, other)
	}

	repays[sortedDebts[length-1].index] = total

	return repays
}
