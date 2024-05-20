// go install golang.org/x/tools/cmd/godoc@latest
// godoc -http :8080
// open http://localhost:8080/

package area

import "sort"

// MinCoins принимает сумму и список монет и возвращает список монет,
// необходимых для составления этой суммы. Функция работает корректно,
// только если список монет отсортирован по убыванию. В противном случае
// может дать неверный результат. Не оптимизирована: использует многократное
// выделение памяти для слайса результатов.

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

// MinCoins2 оптимизированная версия функции MinCoins.
// В начале функция сортирует список монет по убыванию,
// чтобы обеспечить корректность работы. Затем функция использует
// однократное выделение памяти для слайса результатов, что
// предотвращает многократное перевыделение памяти при каждом добавлении элемента.
// Функция также проверяет, не осталась ли неиспользованная сумма,
// что гарантирует корректность результата.

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
