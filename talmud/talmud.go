package talmud

import (
	"sort"

	"github.com/xuender/helper/cont"
	"github.com/xuender/helper/types"
)

func Talmud[N types.Number](total N, debts ...N) ([]N, N) {
	if total == 0 {
		return make([]N, len(debts)), 0
	}

	if len(debts) == 0 {
		return nil, total
	}

	isNegative := total < 0
	if isNegative {
		total = -total
	}

	results, remaining := talmud(total, debts)

	if isNegative {
		for i := range results {
			results[i] = -results[i]
		}

		if remaining != 0 {
			remaining = -remaining
		}
	}

	return results, remaining
}

func talmud[N types.Number](total N, debts []N) ([]N, N) {
	sum := cont.Sum(debts)

	var remaining N

	for total > sum {
		remaining += sum
		total -= sum
	}

	if total == sum {
		return debts, remaining
	}

	sortedDebts := make([]debtWithIndex[N], len(debts))
	for index, debt := range debts {
		sortedDebts[index] = debtWithIndex[N]{debt: debt, index: index}
	}

	sort.Slice(sortedDebts, func(i, j int) bool {
		return sortedDebts[i].debt < sortedDebts[j].debt
	})

	sortedResults := allocate(sortedDebts, total)

	results := make([]N, len(debts))
	for i, s := range sortedDebts {
		results[s.index] = sortedResults[i]
	}

	return results, remaining
}

func allocate[N types.Number](sorted []debtWithIndex[N], remaining N) []N {
	length := len(sorted)
	if length == 1 {
		return []N{min(sorted[0].debt, remaining)}
	}

	var two N = 2
	// 计算分界点
	firstDebt := sorted[0].debt
	// 第一分界点：当前最小债权的一半乘以人数
	eStar := firstDebt * N(length) / two

	var results []N

	if remaining <= eStar {
		// 情况1：总金额低于第一分界点，所有人平分
		each := remaining / N(length)
		results = make([]N, length)

		for index := range results {
			results[index] = each
		}

		return results
	}
	// 情况2：总金额高于第一分界点，递归处理
	// 先给第一个人分配其债权的一半
	firstShare := firstDebt / two
	remainingAfterFirst := remaining - firstShare
	// 剩余金额分配给其余n-1人
	restResults := allocate(sorted[1:], remainingAfterFirst)
	// 组合结果
	results = append([]N{firstShare}, restResults...)
	// 确保每个人的分配不超过其债权（递归过程中可能出现超额）
	for i := range results {
		results[i] = min(results[i], sorted[i].debt)
	}

	return results
}
