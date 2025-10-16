package talmud

import "github.com/xuender/helper/types"

// Dichotomy implements the Talmudic principle of debt repayment allocation using a binary search approach.
// It fairly distributes a total payment amount between two debts according to Talmudic law.
//
// Parameters:
//   - total: The total amount available for debt repayment
//   - debt1: The amount of the first debt
//   - debt2: The amount of the second debt
//
// Returns:
//   - The first return value: Amount to be repaid for the first debt
//   - The second return value: Amount to be repaid for the second debt
func Dichotomy[N types.Number](total, debt1, debt2 N) (N, N) {
	if total == 0 {
		return 0, 0
	}

	isNegative := total < 0
	if isNegative {
		total = -total
	}

	repay1, repay2 := dichotomy(total, debt1, debt2)

	if isNegative {
		return -repay1, -repay2
	}

	return repay1, repay2
}

func dichotomy[N types.Number](total, debt1, debt2 N) (N, N) {
	if total >= debt1+debt2 {
		return debt1, debt2
	}

	swapped := debt1 > debt2
	if swapped {
		debt1, debt2 = debt2, debt1
	}

	if total <= debt1 || debt1 == debt2 {
		half := total / _two

		return half, half
	}

	uncontested := total - debt1
	repay1, repay2 := dichotomy(debt1, debt1, debt2-uncontested)

	if swapped {
		return repay2 + uncontested, repay1
	}

	return repay1, repay2 + uncontested
}
