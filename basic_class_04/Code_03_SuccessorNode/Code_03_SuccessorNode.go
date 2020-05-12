package main

import "fmt"

type Node struct {
	value  int
	left   *Node
	right  *Node
	parent *Node
}

func successorNode(node *Node) *Node {
	if node == nil {
		return node
	}

	if node.right != nil {
		return getLeftMost(node.right)
	} else {
		for node.parent != nil && node.parent.left != nil {
			node = node.parent
			node.parent = node.parent.parent
		}
		return node.parent
	}

}

func getLeftMost(node *Node) *Node {
	if node == nil {
		return node
	}

	for node.left != nil {
		node = node.left
	}
	return node
}

// for test
func main() {
	head := new(Node)
	head.value = 1
	head.parent = nil
	head.left = new(Node)
	head.right = new(Node)

	head.left.value = 2
	head.left.parent = head
	head.left.left = new(Node)
	head.left.right = new(Node)

	head.right.value = 3
	head.right.parent = head
	head.right.left = new(Node)
	head.right.right = new(Node)

	head.left.left.value = 4
	head.left.left.parent = head.left
	head.left.left.left = nil
	head.left.left.right = nil

	head.left.right.value = 5
	head.left.right.parent = head.left
	head.left.right.left = nil
	head.left.right.right = nil

	head.right.left.value = 6
	head.right.left.parent = head.right
	head.right.left.left = nil
	head.right.left.right = nil

	head.right.right.value = 7
	head.right.right.parent = head.right
	head.right.right.left = nil
	head.right.right.right = nil

	test := head
	fmt.Println(test.value, successorNode(test).value)

	test = head.left
	fmt.Println(test.value, successorNode(test).value)

	test = head.right.right
	fmt.Println(test.value, successorNode(test).value)
}
