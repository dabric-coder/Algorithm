package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	value int
	left *Node
	right *Node
}


func serialByPre(head *Node) string {
	if head == nil {
		return "#!"
	}

	var res string
	res = strconv.Itoa(head.value) + "!"
	res += serialByPre(head.left)
	res += serialByPre(head.right)
	return res
}

func reconByPreString(preStr string) *Node {
	values := strings.Split(preStr, "!")
	if values[len(values)-1] == "" {
		values = values[:len(values)-1]
	}
	fmt.Println(values)
	return reconPreOrder(values)
}

func reconPreOrder(values []string) *Node {
	value := values[0]
	values = append(values[:0], values[1:]...)

	fmt.Println(values)
	if value == "#" {
		return nil
	}

	head := new(Node)
	head.value, _ = strconv.Atoi(value)
	head.left = reconPreOrder(values)
	head.right = reconPreOrder(values)
	return head
}

func main() {
	head := new(Node)
	head.value = 1
	head.left = new(Node)
	head.right = new(Node)

	head.left.value = 2
	head.right.value = 3

	head.left.left = new(Node)
	head.left.left.value = 4

	head.right.right = new(Node)
	head.right.right.value = 5

	values := serialByPre(head)
	fmt.Println(serialByPre(head))
	new_head := reconByPreString(values)

	fmt.Println(new_head.value, new_head.left.value,
		new_head.right.value, new_head.left.left.value,
		new_head.right.right.value)


}