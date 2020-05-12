package main

import "fmt"

type Node struct {
	value  int
	next   *Node
	random *Node
}

func copyListWithRand1(head *Node) *Node {
	var nodeMap map[*Node]*Node
	nodeMap = make(map[*Node]*Node, 2)

	cur := head
	for cur != nil {
		copyNode := new(Node)
		copyNode.value = cur.value
		nodeMap[cur] = copyNode
		cur = cur.next
	}

	//for k, v := range nodeMap {
	//	fmt.Printf("%v, %v\n", k, v)
	//}
	//fmt.Println()

	cur = head
	for cur != nil {
		nodeMap[cur].next = nodeMap[cur.next]
		nodeMap[cur].random = nodeMap[cur.random]
		cur = cur.next
	}

	return nodeMap[head]
}

func copyListWithRand2(head *Node) *Node {
	cur := head

	// 将复制的节点插入原节点后面
	for cur != nil {
		next := cur.next
		copyNode := new(Node)
		copyNode.value = cur.value
		cur.next = copyNode
		copyNode.next = next
		cur = next
	}

	// 复制原节点的中random指针指向
	cur = head
	for cur != nil {
		next := cur.next.next
		curCopy := cur.next
		if cur.random != nil {
			curCopy.random = cur.random.next
		} else {
			curCopy.random = nil
		}
		cur = next
	}

	// 将链表复原
	cur = head
	res := head.next

	for cur != nil {
		next := cur.next.next
		curCopy := cur.next
		cur.next = next
		if next != nil {
			curCopy.next = next.next
		} else {
			curCopy.next = nil
		}
		cur = next
	}
	return res
}

func printList(head *Node) {
	if head == nil {
		return
	}

	for head != nil {
		fmt.Printf("%v节点的地址是%v, next指向%v, random指向%v\n", head.value, head, head.next, head.random)
		head = head.next
	}
	fmt.Println()
}

func main() {
	head := new(Node)
	head.value = 1
	head.next = new(Node)
	head.next.value = 2
	head.next.random = nil
	head.next.next = new(Node)
	head.next.next.value = 3
	head.next.next.random = head

	head.random = head.next.next

	printList(head)

	new_head := copyListWithRand2(head)
	printList(new_head)
}
