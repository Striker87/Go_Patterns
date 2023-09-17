package main

import "fmt"

// кольцевой буффер является распространной реализацией очереди
// используется когда нужно успевать обрабатывать входной потом данных где некоторые элементы входного потока могут быть без последствий отброшены
// таким образом прореживая слишком быстро поступающий поток
// соединяем два накала через одну горутину
// применяется например когда пишет метрики и потеря некоторых данных не критична в случае когда читатель входного канала не успевает их обрабатывать

type RightBuffer struct {
	inCh  chan int
	outCh chan int
}

func main() {
	inCh := make(chan int)
	outCh := make(chan int, 4)
	rb := NewRingBuffer(inCh, outCh)

	go rb.Run() // Run() должно быть вызвано один раз

	for i := 1; i <= 10; i++ {
		inCh <- i
	}
	close(inCh)

	for res := range outCh {
		fmt.Println(res)
	}
}

func NewRingBuffer(in, out chan int) *RightBuffer {
	return &RightBuffer{
		inCh:  in,
		outCh: out,
	}
}

func (r *RightBuffer) Run() {
	defer close(r.outCh)

	for v := range r.inCh {
		select {
		case r.outCh <- v:
		default:
			<-r.outCh
			r.outCh <- v
		}
	}
}
