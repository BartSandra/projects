package mincoins

import (
	"reflect"
	"testing"
)

func TestMinCoins2(t *testing.T) {
	tests := []struct {
		name     string
		val      int
		coins    []int
		expected []int
	}{
		{"Test 1", 15, []int{1, 2, 3, 4}, []int{4, 4, 4, 3}},
		{"Test 2", 123, []int{1, 5, 10, 100}, []int{100, 10, 10, 1, 1, 1}},
		{"Test 3", 160, []int{1, 5, 10, 50, 100}, []int{100, 50, 10}},
		{"Test 4", 23, []int{1, 2, 5, 8}, []int{8, 8, 5, 2}},
		{"Test 5", 7, []int{1, 2}, []int{2, 2, 2, 1}},
		{"Test 6", 190, []int{10, 20, 20, 80}, []int{80, 80, 20, 10}},
		{"Test 7", 16, []int{7, 5, 2, 1}, []int{7, 7, 2}},
		{"Test 8", 0, []int{7, 5, 2, 1}, []int{}},
		{"Test 9", 30, []int{10, 20, 40}, []int{20, 10}},
		{"Test 10", 15, []int{10, 20, 40}, []int{}},
		{"Test 11", 13, []int{1, 5, 10}, []int{10, 1, 1, 1}},
		{"Test 12", 0, []int{1, 5, 10}, []int{}},
		{"Test 13", 15, []int{10, 20, 40}, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MinCoins2(tt.val, tt.coins)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("%s: MinCoins2(%d, %v) = %v, want %v", tt.name, tt.val, tt.coins, got, tt.expected)
			}
		})
	}
}

func TestMinCoins(t *testing.T) {
	tests := []struct {
		name     string
		val      int
		coins    []int
		expected []int
	}{
		{"Test 1", 15, []int{1, 2, 3, 4}, []int{4, 4, 4, 3}},
		{"Test 2", 123, []int{1, 5, 10, 100}, []int{100, 10, 10, 1, 1, 1}},
		{"Test 3", 160, []int{1, 5, 10, 50, 100}, []int{100, 50, 10}},
		{"Test 4", 23, []int{1, 2, 5, 8}, []int{8, 8, 5, 2}},
		{"Test 5", 7, []int{1, 2}, []int{2, 2, 2, 1}},
		{"Test 6", 190, []int{10, 20, 20, 80}, []int{80, 80, 20, 10}},
		{"Test 8", 0, []int{1, 2, 3, 4}, []int{}},
		{"Test 9", 30, []int{10, 20, 40}, []int{20, 10}},
		{"Test 10", 13, []int{1, 5, 10}, []int{10, 1, 1, 1}},
		//{"Test 11", 15, []int{10, 20, 40}, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MinCoins(tt.val, tt.coins)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("%s: MinCoins(%d, %v) = %v, want %v", tt.name, tt.val, tt.coins, got, tt.expected)
			}
		})
	}
}
