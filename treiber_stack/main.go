package main

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

// lock free stack

type item struct {
	value int
	next  unsafe.Pointer
}

type Stack struct {
	head unsafe.Pointer
}

func NewStack() Stack {
	return Stack{}
}

func (s *Stack) Push(value int) {
	node := &item{value: value} // создаем новую ноду

	for {
		head := atomic.LoadPointer(&s.head) // загружаем значения предудущей головы связанного списка
		node.next = head

		if atomic.CompareAndSwapPointer(&s.head, head, unsafe.Pointer(node)) {
			return
		}
	}
}

func (s *Stack) Pop() int {
	for {
		head := atomic.LoadPointer(&s.head)
		if head == nil {
			return -1
		}

		next := atomic.LoadPointer(&(*item)(head).next)
		if atomic.CompareAndSwapPointer(&s.head, head, next) {
			return (*item)(head).value
		}
	}
}

func main() {
	stack := NewStack()

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
