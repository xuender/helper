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
//   - debts: a slice of debt amounts to be repaid
//
// Returns:
//   - A slice of numbers representing the actual repayment amounts for each debt,
//     corresponding to the order of the input debts
func Talmud[N types.Number](total N, debts []N) []N {
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

	isNegative := total < 0
	if isNegative {
		total = -total
	}

	sortedDebts := make([]debtWithIndex[N], length)
	for index, debt := range debts {
		sortedDebts[index] = debtWithIndex[N]{debt: debt, index: index}
	}

	sort.Slice(sortedDebts, func(i, j int) bool {
		return sortedDebts[i].debt < sortedDebts[j].debt
	})

	repays := make([]N, length)

	if total <= sortedDebts[0].debt/2*N(length) {
		return applySign(isNegative, repayEach(total, length, repays))
	}

	for index, current := range sortedDebts[:length-1] {
		var totalOther N

		for _, debt := range sortedDebts[index+1:] {
			totalOther += debt.debt
		}

		repays[current.index], total = Dichotomy(total, current.debt, totalOther)
	}

	repays[sortedDebts[length-1].index] = total

	return applySign(isNegative, repays)
}

func repayEach[N types.Number](total N, length int, repays []N) []N {
	each := total / N(length)

	for i := range length {
		repays[i] = each
	}

	return repays
}

func applySign[N types.Number](isNegative bool, repays []N) []N {
	if isNegative {
		for index := range repays {
			repays[index] = -repays[index]
		}
	}

	return repays
}
