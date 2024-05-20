package main

import (
	"testing"
)

func TestMultiplex(t *testing.T) {
	c1 := make(chan interface{})
	c2 := make(chan interface{})
	c3 := make(chan interface{})
	c4 := make(chan interface{})
	c5 := make(chan interface{})

	go func() { defer close(c1); c1 <- "Hello" }()
	go func() { defer close(c2); c2 <- 42 }()
	go func() { defer close(c3); c3 <- "Sandra" }()
	go func() { defer close(c4); c4 <- 777 }()
	go func() {
		defer close(c5)
		for _, x := range []int{1, 2, 3} {
			c5 <- x
		}
	}()

	expectedResults := map[interface{}]bool{
		"Hello":  true,
		42:       true,
		"Sandra": true,
		777:      true,
		1:        true,
		2:        true,
		3:        true,
	}

	multiplexed := multiplex(c1, c2, c3, c4, c5)

	for val := range multiplexed {
		if _, ok := expectedResults[val]; !ok {
			t.Errorf("Unexpected value: %v", val)
		} else {
			delete(expectedResults, val)
		}
	}

	if len(expectedResults) != 0 {
		t.Errorf("Not all values received, missing: %v", expectedResults)
	}
}
