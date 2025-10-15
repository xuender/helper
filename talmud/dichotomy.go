package talmud

import "github.com/xuender/helper/types"

// Dichotomy implements the Talmudic principle of debt repayment allocation using a binary search approach.
// It fairly distributes a total payment amount between two debts according to Talmudic law.
//
// Parameters:
//   total: The total amount available for debt repayment
//   debt1: The amount of the first debt
//   debt2: The amount of the second debt
//
// Returns:
//   The first return value: Amount to be repaid for the first debt
//   The second return value: Amount to be repaid for the second debt

func Dichotomy[N types.Number](total, debt1, debt2 N) (N, N) {
	if total >= debt1+debt2 {
		return debt1, debt2
	}

	if debt1 > debt2 {
		repay1, repay2 := Dichotomy(total, debt2, debt1)

		return repay2, repay1
	}

	if total <= debt1 || debt1 == debt2 {
		haft := total / _two

		return haft, haft
	}

	repay := total - debt1
	repay1, repay2 := Dichotomy(debt1, debt1, debt2-repay)

	return repay1, repay2 + repay
}
