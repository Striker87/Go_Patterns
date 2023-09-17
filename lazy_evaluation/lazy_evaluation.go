package main

import (
	"fmt"
	"sync"
)

type LazyInt func() int

// ленивые вычисления
func main() {
	n := makeLazy(func() int {
		fmt.Println("Doing expensive calculation")
		return 23
	})

	fmt.Println(n())      // calculates 23
	fmt.Println(n() + 42) // reuses the calculated value
}

func makeLazy(f LazyInt) LazyInt {
	var v int
	var once sync.Once
	return func() int {
		once.Do(func() {
			v = f()
			f = nil // пересохраняем чтоб сборщик мусора ее удалил
		})
		return v
	}
}
