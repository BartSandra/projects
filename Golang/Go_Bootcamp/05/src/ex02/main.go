/*принимает число n и кучу подарков. Она возвращает слайс из n “самых крутых” подарков,
основываясь на их значениях и размерах. Если запрошено больше подарков, чем доступно,
cdфункция возвращает ошибку.*/

package main

import (
	"container/heap"
	"fmt"
)

type Present struct {
	Value int
	Size  int
}

type PresentHeap []Present

func (p PresentHeap) Len() int {
	return len(p)
}

// определяем порядок элементов
func (p PresentHeap) Less(i, j int) bool {
	if p[i].Value == p[j].Value {
		return p[i].Size < p[j].Size
	}
	return p[i].Value > p[j].Value
}

func (p PresentHeap) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PresentHeap) Push(x interface{}) {
	item := x.(Present)
	*p = append(*p, item)
}

func (p *PresentHeap) Pop() interface{} {
	old := *p
	n := len(old)
	item := old[n-1]
	*p = old[0 : n-1]
	return item
}

func NCoolestPresents(n int, presents *PresentHeap) (*PresentHeap, error) {
	if n > presents.Len() {
		return nil, fmt.Errorf("ошибка: количество запрошенных подарков больше количества доступных подарков")
	}

	heap.Init(presents)

	res := make(PresentHeap, 0, n)
	for i := 0; i < n; i++ {
		res = append(res, heap.Pop(presents).(Present))
	}

	return &res, nil
}

func main() {
	presents := &PresentHeap{
		{Value: 5, Size: 1},
		{Value: 4, Size: 5},
		{Value: 3, Size: 1},
		{Value: 5, Size: 2},
	}

	n := 2

	coolestPresents, err := NCoolestPresents(n, presents)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Println(*coolestPresents)
}
