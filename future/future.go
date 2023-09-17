package main

import (
	"fmt"
	"time"
)

type data struct {
	Body  string
	Error error
}

// позволяет запустить вычисление некоторых данных в фоне не дожидаясь их обработки
// в случае если мы знаем что эти данные понадобятся нам в дальнейшем, но не прямо сейчас
func main() {
	future := future("https://i.ua")
	fmt.Println("request started")
	body := <-future

	fmt.Printf("Response: %s", body)
}

func future(link string) <-chan data {
	ch := make(chan data, 1)

	go func() {
		body, err := doGet(link)

		ch <- data{Body: body, Error: err}
	}()

	return ch
}

func doGet(link string) (string, error) {
	// parsing url
	time.Sleep(time.Millisecond * 200)
	return fmt.Sprintf("Resp: %s", link), nil
}
