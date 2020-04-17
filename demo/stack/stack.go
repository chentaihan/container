package main

import (
	"fmt"
	"github.com/chentaihan/container/stack"
)

func main() {
	var s stack.IStack = stack.NewStack(100)
	//var s stack.IStack = stack.NewStackLink()
	size := 10
	for i := 0; i < size; i++ {
		s.Push(i)
	}
	if s.Len() != size {
		fmt.Println("len error")
	}
	fmt.Println("cap=", s.Cap())
	for !s.Empty() {
		if val, exist := s.Top(); exist {
			fmt.Println(val)
		}
		s.Pop()
	}
	if s.Len() != 0 {
		fmt.Println("empty error")
	}
}
