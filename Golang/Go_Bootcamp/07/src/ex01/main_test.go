//go test -bench=. -cpuprofile=cpu.prof
//go tool pprof cpu.prof
//top10 > top10.txt

package mincoins

import "testing"

var tests = map[string]struct {
	name     string
	coins    []int
	val      int
	expected []int
}{
	"Test 1":  {coins: []int{1, 2, 3, 4}, val: 15, expected: []int{4, 4, 4, 3}},
	"Test 2":  {coins: []int{1, 5, 10, 100}, val: 123, expected: []int{100, 10, 10, 1, 1, 1}},
	"Test 3":  {coins: []int{1, 5, 10, 50, 100}, val: 160, expected: []int{100, 50, 10}},
	"Test 4":  {coins: []int{1, 2, 5, 8}, val: 23, expected: []int{8, 8, 5, 2}},
	"Test 5":  {coins: []int{1, 2}, val: 7, expected: []int{2, 2, 2, 1}},
	"Test 6":  {coins: []int{10, 20, 20, 80}, val: 190, expected: []int{80, 80, 20, 10}},
	"Test 7":  {coins: []int{1, 2, 3, 4}, val: 0, expected: []int{}},
	"Test 8":  {coins: []int{10, 20, 40}, val: 30, expected: []int{20, 10}},
	"Test 9":  {coins: []int{10, 20, 40}, val: 15, expected: []int{}},
	"Test 10": {coins: []int{1, 5, 10}, val: 3, expected: []int{10, 1, 1, 1}},
	"Test 11": {coins: []int{1, 5, 10}, val: 0, expected: []int{}},
}

func BenchmarkMinCoins(b *testing.B) {
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				MinCoins(tt.val, tt.coins)
			}
		})
	}
}

func BenchmarkMinCoins2(b *testing.B) {
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				MinCoins2(tt.val, tt.coins)
			}
		})
	}
}
