package main

import (
	"fmt"
	"github.com/emirpasic/gods/stacks/arraystack"
)


type specialStack struct {
	dataStack *arraystack.Stack
	minStack *arraystack.Stack
}


func (s *specialStack) push(value int) {
	if s.dataStack.Empty() {
		s.dataStack.Push(value)
		s.minStack.Push(value)
	} else {
		min, _ := s.minStack.Peek()
		if value < min.(int) {
			s.dataStack.Push(value)
			s.minStack.Push(value)
		} else {
			s.dataStack.Push(value)
			s.minStack.Push(min)
		}
	}
}

func (s *specialStack) pop() (value int) {
	v, ok := s.dataStack.Pop()
	s.minStack.Pop()
	if ok {
		value = v.(int)
	}
	return
}

func (s *specialStack) getMin() (value int) {
	v, ok := s.minStack.Peek()
	if ok {
		value = v.(int)
	}
	return
}

func main() {
	var s = specialStack{arraystack.New(), arraystack.New()}
	fmt.Println(s.getMin())
	fmt.Println(s.getMin())
	s.push(4)
	fmt.Println(s.getMin())
	s.push(5)
	fmt.Println(s.getMin())
	s.push(3)
	fmt.Println(s.getMin())
	s.pop()
	fmt.Println(s.getMin())


}
