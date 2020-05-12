package arraystack

import (
	"bytes"
	"fmt"
)

type Stack struct {
	array []interface{}
	size  int
}

func newStack(capacity int) *Stack {
	return &Stack{make([]interface{}, capacity), 0}
}

func (s *Stack) push(e interface{}) error {
	if s.size > cap(s.array) {
		err := fmt.Errorf("The stack is full.")
		return err
	}
	s.array[s.size] = e
	s.size++
	return nil
}

func (s *Stack) pop() (interface{}, error) {
	if s.isEmpty() {
		err := fmt.Errorf("the stack is null")
		return nil, err
	}
	data := s.array[s.size-1]

	s.array = s.array[:s.size-1]
	s.size--
	return data, nil
}

func (s *Stack) peek() interface{} {
	if s.isEmpty() {
		return nil
	}

	return s.array[s.size-1]
}

func (s *Stack) isEmpty() bool {
	if s.size == 0 {
		return true
	}
	return false

}

func (s *Stack) getSize() int {
	return s.size
}

func (s *Stack) printStack() {
	fmt.Printf("stack: [")
	for i := 0; i < s.size; i++ {
		fmt.Printf("%d ", s.array[i])
	}
	fmt.Println("] top")
}

// 重写栈打印时的展示形式，只需要重新String方法
func (s *Stack) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf(""))
	fmt.Println()

}
