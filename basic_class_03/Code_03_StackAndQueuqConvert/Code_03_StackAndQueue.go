package main

import (
	"fmt"
	"github.com/emirpasic/gods/stacks/arraystack"
	"log"
)

type stackToQueue struct {
	pushStack *arraystack.Stack
	popStack *arraystack.Stack
}

func (q *stackToQueue) push(value int) {
	q.pushStack.Push(value)
}

func (q *stackToQueue) isEmpty() bool {
	if q.pushStack.Empty() {
		return true
	}
	return false
}

func (q *stackToQueue) pour() {
	if q.pushStack.Empty() && q.popStack.Empty() {
		log.Fatal("The queue is empty")
	} else if q.popStack.Empty() {
		for !q.pushStack.Empty() {
			v, _ := q.pushStack.Pop()
			q.popStack.Push(v)
		}
	}
}

func (q *stackToQueue) poll() (value int) {
	q.pour()
	v, ok := q.popStack.Pop()
	if ok {
		value = v.(int)
	}
	return
}

func (q *stackToQueue) size() int {
	return q.pushStack.Size()
}

func (q *stackToQueue) peek() (value int) {
	q.pour()
	v, ok := q.popStack.Peek()
	if ok {
		value = v.(int)
	}
	return
}

// 在上面的栈实现的队列上再实现栈
type queueToStack struct {
	queue *stackToQueue
	help *stackToQueue
}

func (s *queueToStack) push(value int) {
	s.queue.push(value)
}

func (s *queueToStack) pop() int {
	if s.queue.isEmpty() {
		log.Fatal("stack is empty.")
	}

	for s.queue.size() != 1 {
		s.help.push(s.queue.poll())
	}

	res := s.queue.poll()
	s.swap()
	return res
}

func (s *queueToStack) peek() int {
	if s.queue.isEmpty() {
		log.Fatal("stack is empty.")
	}

	for s.queue.size() != 1 {
		s.help.push(s.queue.poll())
	}

	res := s.queue.poll()
	s.help.push(res)
	s.swap()
	return res
}

func (s *queueToStack) swap() {
	tmp := s.queue
	s.queue = s.help
	s.help = tmp
}


func main() {
	var queue = stackToQueue{arraystack.New(), arraystack.New()}
	var help = stackToQueue{arraystack.New(), arraystack.New()}

	var stack =  queueToStack{&queue, &help}

	stack.push(1)
	stack.push(2)
	stack.push(3)
	stack.push(4)

	//fmt.Println(stack.pop())
	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
	fmt.Println(stack.pop())





}