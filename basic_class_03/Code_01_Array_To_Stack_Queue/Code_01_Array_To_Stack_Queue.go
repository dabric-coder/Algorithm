package main

import (
	"fmt"
	"math"
)

type arrayStack struct {
	arr   []int
	index int
}

func (s *arrayStack) push(value int) error {
	if s.index == len(s.arr) {
		err := fmt.Errorf("栈满了")
		return err
	}
	s.arr[s.index] = value
	s.index++
	return nil
}

func (s *arrayStack) pop() (int, error) {
	if s.index == 0 {
		err := fmt.Errorf("栈空了")
		return math.MaxInt64, err
	}
	s.index--
	return s.arr[s.index], nil
}

func (s *arrayStack) peek() int {
	if s.index == 0 {
		return math.MaxInt64
	}

	return s.arr[s.index-1 ]
}

// 工厂函数，初始化arrayStack
func initStack(initSize int) (*arrayStack, error) {

	if initSize < 0 {
		err := fmt.Errorf("栈的大小不能小于0")
		return nil, err
	}

	stack := &arrayStack{
		arr:   make([]int, initSize),
		index: 0,
	}
	return stack, nil
}

type arrayQueue struct {
	arr   []int
	size  int
	end   int
	start int
}


func (q *arrayQueue) push(value int) error {
	if q.size == len(q.arr) {
		err := fmt.Errorf("队列满了")
		return err
	}

	q.size++
	q.arr[q.end] = value


	if q.end == len(q.arr) - 1 {
		q.end = 0
	} else {
		q.end++
	}
	return nil
}

func (q *arrayQueue) poll() (int, error) {
	if q.size == 0 {
		err := fmt.Errorf("栈空了")
		return math.MaxInt64, err
	}

	q.size--
	tmp := q.start
	if q.start == len(q.arr) - 1 {
		q.start = 0
	} else {
		q.start++
	}
	return q.arr[tmp], nil
}

func (q *arrayQueue) peek() int {
	if q.size == 0 {
		return math.MaxInt64
	}

	return q.arr[q.start]
}

func initQueue(initSize int) (*arrayQueue, error) {
	if initSize < 0 {
		err := fmt.Errorf("栈的大小不能小于0")
		return nil, err
	}

	queue := &arrayQueue{
		arr:   make([]int, initSize),
		size:  0,
		end:   0,
		start: 0,
	}

	return queue, nil
}

func main() {
	queue, _ := initQueue(3)

	for i := 1; i < 4; i++ {
		err := queue.push(i)
		if err != nil {
			fmt.Println(err)
		}
	}


	fmt.Println(queue.peek())
	fmt.Println(queue.poll())

	fmt.Println(queue.peek())

	fmt.Println(queue.poll())
	fmt.Println(queue.poll())
	fmt.Println(queue.poll())
	fmt.Println(queue.poll())
}
