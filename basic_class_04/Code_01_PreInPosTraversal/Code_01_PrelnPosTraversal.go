package main

import (
	"fmt"
	"github.com/emirpasic/gods/stacks/arraystack"
)

type Node struct {
	left *Node
	right *Node
	value int
}

func preOrderRecur(head *Node) {
	if head == nil {
		return
	}
	fmt.Printf("%v ", head.value)
	preOrderRecur(head.left)
	preOrderRecur(head.right)
}
func inOrderRecur(head *Node) {
	if head == nil {
		return
	}
	inOrderRecur(head.left)
	fmt.Printf("%v ", head.value)
	inOrderRecur(head.right)
}

func posOrderRecur(head *Node) {
	if head == nil {
		return
	}
	posOrderRecur(head.left)
	posOrderRecur(head.right)
	fmt.Printf("%v ", head.value)
}

func preOrderUnrecur(head *Node) {
	if head != nil {
		stack := arraystack.New()
		stack.Push(head)
		for !stack.Empty() {
			cur, _ := stack.Pop()
			fmt.Printf("%v ", cur.(*Node).value)
			if cur.(*Node).right != nil {
				stack.Push(cur.(*Node).right)
			}
			if cur.(*Node).left != nil {
				stack.Push(cur.(*Node).left)
			}
		}
	}
	fmt.Println()
}

func inOrderUnrecur(head *Node) {
	if head != nil {
		stack := arraystack.New()
		cur := head
		for !stack.Empty() || cur != nil {
			if cur != nil {
				stack.Push(cur)
				cur = cur.left
			} else {
				tmp, _ := stack.Pop()
				cur = tmp.(*Node)
				fmt.Printf("%v ", cur.value)
				cur = cur.right
			}
		}
	}
	fmt.Println()
}


func posOrderUnrecur(head *Node) {
	if head != nil {
		s1 := arraystack.New()
		s2 := arraystack.New()
		s1.Push(head)
		for !s1.Empty() {
			cur, _ := s1.Pop()
			s2.Push(cur)
			if cur.(*Node).left != nil {
				s1.Push(cur.(*Node).left)
			}
			if cur.(*Node).right != nil {
				s1.Push(cur.(*Node).right)
			}
		}
		for !s2.Empty() {
			tmp, _ := s2.Pop()
			fmt.Printf("%v ", tmp.(*Node).value)
		}
	}
	fmt.Println()
}

func main() {
	head := new(Node)
	head.value = 1

	head.left = new(Node)
	head.left.value = 2
	head.right = new(Node)
	head.right.value = 3

	head.left.left = new(Node)
	head.left.left.value = 4
	head.left.right = new(Node)
	head.left.right.value = 5

	head.right.left = new(Node)
	head.right.left.value = 6
	head.right.right = new(Node)
	head.right.right.value = 7

	preOrderRecur(head)
	fmt.Println()

	inOrderRecur(head)
	fmt.Println()

	posOrderRecur(head)
	fmt.Println()
	fmt.Println("=====================")
	preOrderUnrecur(head)
	inOrderUnrecur(head)
	posOrderUnrecur(head)
}