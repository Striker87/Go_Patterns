package main

import (
	"fmt"
	"sync"
	"time"
)

// комбинирует несколько входов в один выходной канал (мультиплексируем канал)
// порядок вывода не гарантируется

type Payload struct {
	Name  string
	Value int
}

func main() {
	done := make(chan struct{})
	wg := sync.WaitGroup{}

	wg.Add(2)
	producers := make([]<-chan Payload, 0, 3)
	producers = append(producers, producer("Alice", done, &wg))
	producers = append(producers, producer("Jack", done, &wg))

	fanIn := make(chan Payload)

	wg.Add(2)
	consumer(producers, done, &wg, fanIn)

	go func() {
		for {
			select {
			case <-done:
				return
			case v := <-fanIn:
				fmt.Printf("fanIn got %v\n", v)
			}
		}
	}()

	time.Sleep(time.Second)
	close(done)
	wg.Wait()
}

func producer(name string, done <-chan struct{}, wg *sync.WaitGroup) <-chan Payload {
	ch := make(chan Payload)
	var i = 1

	go func() {
		defer wg.Done()

		for {
			select {
			case <-done:
				close(ch)
				fmt.Println(name, " completed")
				return
			case ch <- Payload{
				Name:  name,
				Value: i,
			}:
				fmt.Println(name, " produced", i)
				i++
				time.Sleep(time.Millisecond * 500)
			}
		}
	}()

	return ch
}

func consumer(channels []<-chan Payload, done <-chan struct{}, wg *sync.WaitGroup, fanIn chan<- Payload) {
	for i, ch := range channels {
		i := i + 1
		ch := ch

		go func() {
			defer wg.Done()
			fmt.Println("started consume of producer", i)

			for {
				select {
				case <-done:
					fmt.Printf("consume of producer %d completed", i)
					return
				case v := <-ch:
					fmt.Printf("consume of producer %d got value %d from %s\n", i, v.Value, v.Name)
					fanIn <- v
				}
			}
		}()
	}
}
