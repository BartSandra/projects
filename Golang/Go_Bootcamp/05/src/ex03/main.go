/*В данном случае, рюкзаком является набор подарков (PresentHeap), каждый из которых имеет
определенное значение (Value) и размер (Size).

Функция GrabPresents принимает вместимость рюкзака и набор подарков, и возвращает максимальное значение,
которое можно уложить в рюкзак данной вместимости.*/

package main

import (
	"fmt"
	"log"
)

type Present struct {
	Value int
	Size  int
}

type PresentHeap []Present

func (p PresentHeap) Len() int {
	return len(p)
}

func GrabPresents(cap int, p *PresentHeap) (int, error) {
	if cap < 0 {
		return 0, fmt.Errorf("некорректная вместимость")
	}

	n := p.Len()
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, cap+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= cap; j++ {
			if (*p)[i-1].Size > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-(*p)[i-1].Size]+(*p)[i-1].Value)
			}
		}
	}

	return dp[n][cap], nil
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	p := &PresentHeap{
		{Value: 3, Size: 5},
		{Value: 5, Size: 10},
		{Value: 4, Size: 6},
		{Value: 2, Size: 5},
	}

	result, err := GrabPresents(14, p)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(result)
}
