package main

import "fmt"

// стек на основе связного списка

type item struct {
	value int   // значение
	next  *item // указатель на следующий элемент
}

type Stack struct { // струкрута стека
	head *item
}

func NewStack() Stack {
	return Stack{}
}

func (s *Stack) Push(value int) { // создаем ноду, присваиваем следующее значение головы
	s.head = &item{
		value: value,
		next:  s.head,
	}
}

func (s *Stack) Pop() int {
	if s.head == nil {
		return -1
	}

	value := s.head.value // запоминаем значение головы
	s.head = s.head.next  // двигаем голову на следующий элемент
	return value
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
