package main

import (
	"fmt"
	"sync"
	"time"
)

// очередь с лимитом на параллелизм

const (
	dataSize       = 10
	semaphoreLimit = 4
)

func main() {
	data := make([]int, 0, dataSize)

	for i := 0; i < dataSize; i++ {
		data = append(data, i+1)
	}

	semaphore := make(chan struct{}, semaphoreLimit)
	results := make([]int, dataSize)

	var wg sync.WaitGroup

	fmt.Printf("Before: %v\n", data)
	start := time.Now()

	for i, xi := range data {
		wg.Add(1)
		go func(i int, xi int) {
			defer wg.Done()
			semaphore <- struct{}{}
			results[i] = calculate(xi)
			<-semaphore
		}(i, xi)
	}

	wg.Wait()
	fmt.Printf("After: %v\n", results)
	fmt.Printf("Elapsed: %s\n", time.Since(start))
}

func calculate(val int) int {
	time.Sleep(time.Millisecond * 500)
	return val * 2
}
