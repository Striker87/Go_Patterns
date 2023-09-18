package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	n        = 3
	messages = 10
)

func main() {
	var wg sync.WaitGroup

	fmt.Println("Queue of length n: ", n)
	queue := make(chan struct{}, n)

	wg.Add(messages)

	for w := 1; w <= messages; w++ {
		process(w, queue, &wg)
	}

	wg.Wait()

	close(queue)

	fmt.Println("processing complete")
}

func process(payload int, queue chan struct{}, wg *sync.WaitGroup) {
	queue <- struct{}{}

	go func() {
		defer wg.Done()

		fmt.Printf("Start processing of %d\n", payload)
		time.Sleep(time.Millisecond * 500)
		fmt.Printf("Complete processing of %d\n", payload)

		<-queue
	}()
}
