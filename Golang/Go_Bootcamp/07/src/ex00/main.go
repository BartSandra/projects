package mincoins

import "sort"

func MinCoins(val int, coins []int) []int {
	res := make([]int, 0)
	i := len(coins) - 1
	for i >= 0 {
		for val >= coins[i] {
			val -= coins[i]
			res = append(res, coins[i])
		}
		i -= 1
	}
	return res
}

func MinCoins2(val int, coins []int) []int {
	sort.Slice(coins, func(i, j int) bool {
		return coins[i] > coins[j]
	})

	res := make([]int, 0)
	for _, coin := range coins {
		for val >= coin {
			val -= coin
			res = append(res, coin)
		}
	}

	if val > 0 {
		return []int{}
	}

	return res
}
