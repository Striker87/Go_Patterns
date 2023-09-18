package main

import (
	"fmt"
	"time"
)

const dataSize = 4

func main() {
	data := make([]int, 0, dataSize)

	for i := 0; i < dataSize; i++ {
		data = append(data, i+10)
	}

	//results := make([]int, dataSize)
	semaphore := make(chan int, dataSize)

	fmt.Printf("Before: %v\n", data)
	start := time.Now()

	for i, xi := range data {
		go func(i int, xi int) {
			semaphore <- calculate(xi)
		}(i, xi)
	}

	for i := 0; i < dataSize; i++ {
		fmt.Printf("calculation completed: %d\n", <-semaphore)
	}

	fmt.Printf("Elapsed: %s\n", time.Since(start))
}

func calculate(val int) int {
	time.Sleep(time.Millisecond * 500)
	return val * 2
}
