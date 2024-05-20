package main

import (
	"fmt"
	"sync"
)

func multiplex(channels ...chan interface{}) chan interface{} {
	var wg sync.WaitGroup
	multiplexedStream := make(chan interface{})

	multiplexer := func(c chan interface{}) {
		defer wg.Done()
		for i := range c {
			multiplexedStream <- i
		}
	}

	wg.Add(len(channels))
	for _, c := range channels {
		go multiplexer(c)
	}

	go func() {
		wg.Wait()
		close(multiplexedStream)
	}()

	return multiplexedStream
}

func main() {
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

	for val := range multiplex(c1, c2, c3, c4, c5) {
		fmt.Println(val)
	}
}
