package main

import (
	"fmt"
	"math"
)

// представляет собой последовательность этапов соединенных каналами где каждый этап представляет собой группу горутин выполняющих одну и туже функцию
// на каждом этапе горутины получают значенияпредыдущего этапа с входных каналов, выполняют их обработку, отправляют на следующий этап новые значение через выходные каналы
// на каждом этапе может быть произвольное количество входных и выходных каналов
// исключение составяют только первый и последний этап где есть исходящие и входные каналы соответственно
// первая стадия продюсер, последняя консьюмер, промежуточные стадии middleware
// используется: к примеру надо обработать данные используя внешние сервисы при этом порядок выполнения строго последовательный
func main() {
	work := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	in := generateWork(work)

	out := filterOdd(in) // оставляем только четные числа
	out = square(out)    // возводим в квадрат
	out = half(out)      // делим пополам

	for value := range out {
		fmt.Println(value)
	}
}

func filterOdd(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for i := range in {
			if i%2 == 0 {
				out <- i
			}
		}
	}()

	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for i := range in {
			value := math.Pow(float64(i), 2)
			out <- int(value)
		}
	}()

	return out
}

func half(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for i := range in {
			value := i / 2
			out <- value
		}
	}()

	return out
}

func generateWork(work []int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		for _, w := range work {
			ch <- w
		}
	}()

	return ch
}
