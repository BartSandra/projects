package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go Listener(cancel)

	data := make(chan string)
	parse := crawlWeb(ctx, data)

	go func(data chan string) {
		defer close(data)
		urls := []string{
			"https://example.com",
			"https://ya.ru",
			"https://golang.org",
		}
		for _, url := range urls {
			data <- url
		}
	}(data)

	counter := 0
	for result := range parse {
		fmt.Printf("Result %d: %s\n", counter, result[:500])
		counter++
	}

	fmt.Println("end...")
}

func crawlWeb(ctx context.Context, urls chan string) chan string {
	data := make(chan string, 8)

	go func(data chan string) {
		defer close(data)
		var wg sync.WaitGroup
		semaphore := make(chan struct{}, 8)

		for url := range urls {
			select {
			case <-ctx.Done():
				return
			case semaphore <- struct{}{}:
				wg.Add(1)
				go func(url string) {
					defer wg.Done()
					defer func() { <-semaphore }()
					body, err := PageParser(ctx, url)
					if err != nil {
						fmt.Printf("Error fetching URL %s: %v\n", url, err)
						return
					}
					select {
					case data <- body:
					case <-ctx.Done():
					}
				}(url)
			}
		}

		wg.Wait()
	}(data)

	return data
}

func Listener(cancel context.CancelFunc) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	fmt.Println("\nReceived shutdown signal, exiting...")
	cancel()
}

func PageParser(ctx context.Context, url string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
