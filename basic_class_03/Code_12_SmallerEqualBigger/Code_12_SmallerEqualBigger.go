package main

import "fmt"

type Node struct {
	value int
	next *Node
}


func listPartition1(head *Node, pivot int) *Node {
	if head == nil {
		return head
	}

	cur := head

	var nodeArr []*Node

	for cur != nil {
		nodeArr = append(nodeArr, cur)
		cur = cur.next
	}
	arrPartition(nodeArr, pivot)
	i := 1
	for ; i < len(nodeArr); i++ {
		nodeArr[i-1].next = nodeArr[i]
	}
	nodeArr[i-1].next = nil
	return nodeArr[0]
}

func arrPartition(nodeArr []*Node, pivot int) {
	less := -1
	more := len(nodeArr)
	index := 0
	for index != more {
		if nodeArr[index].value < pivot {
			swap(nodeArr, less+1, index)
			less++
			index++
		} else if nodeArr[index].value == pivot {
			index++
		} else {
			swap(nodeArr, more-1, index)
			more--
		}
	}
}

func listPartition2(head *Node, pivot int) *Node {
	if head == nil {
		return head
	}

	var (
		sH *Node
		sT *Node
		eH *Node
		eT *Node
		bH *Node
		bT *Node
		next *Node  // 保存下一个节点
	)
	for head != nil {
		next = head.next
		head.next = nil
		if head.value < pivot {
			if sH == nil {
				sH = head
				sT = head
			} else {
				sT.next = head
				sT = head
			}
		} else if head.value == pivot {
			if eH == nil {
				eH = head
				sT = head
			} else {
				eT.next = head
				eT = head
			}
		} else {
			if bH == nil {
				bH = head
				bT = head
			} else {
				bT.next = head
				bT = head
			}
		}
		head = next
	}

	if sT != nil {
		sT.next = eH
		if eT != nil{
			eT = eT
		} else {
			eT = sT
		}
	}

	if eT != nil {
		eT.next = bH
	}

	if sH != nil {
		return  sH
	} else if eH != nil {
		return eH
	} else if bH != nil {
		return bH
	}
	return nil
}


func swap (nodeArr []*Node, x, y int) {
	nodeArr[x], nodeArr[y] = nodeArr[y], nodeArr[x]
}

func printNodeList(head *Node) {
	if head == nil {
		return
	}

	for head != nil {
		fmt.Printf("%v ->", head.value)
		head = head.next
	}
	fmt.Println()
}

func main() {
	head := new(Node)
	head.value = 9
	head.next = new(Node)
	head.next.value = 0
	head.next.next = new(Node)
	head.next.next.value = 4
	head.next.next.next = new(Node)
	head.next.next.next.value = 5
	head.next.next.next.next = new(Node)
	head.next.next.next.next.value = 1
	head.next.next.next.next.next = nil

	printNodeList(head)
	new_head := listPartition2(head, 3)
	printNodeList(new_head)

}