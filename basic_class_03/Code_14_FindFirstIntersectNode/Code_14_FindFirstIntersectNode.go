package main

import (
	"fmt"
	"math"
)

type Node struct {
	Val int
	Next *Node
}

func getLoopNode1(head *Node) *Node  {
	if head == nil {
		return nil
	}
	var nodeMap = make(map[*Node]struct{}, 10)
	cur := head
	for {
		if cur == nil {
			return nil
		}
		if !findRes(cur, nodeMap) {
			nodeMap[cur] = struct{}{}
			cur = cur.Next
		} else {
			break
		}
	}
	return cur
}

// 判断一个map中指定的键是否存在
func findRes(node *Node, nodeMap map[*Node]struct{}) bool {
	_, ok := nodeMap[node]
	if ok {
		return true
	} else {
		return false
	}
}

func getLoopNode(head *Node) *Node {
	if head == nil {
		return nil
	}
	fast := head
	slow := head
	for {
		if fast.Next == nil || fast.Next.Next == nil  {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}

	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}

func noLoop1(head1, head2 *Node) *Node {
	if head1 == nil || head2 == nil  {
		return nil
	}
	var nodeMap = make(map[*Node]struct{}, 10)
	// 将第一个链表的节点存放在nodeMap中
	cur := head1
	for cur != nil {
		nodeMap[cur] = struct{}{}
		cur = cur.Next
	}

	// 遍历第二个链表，看是否有节点在nodeMap中
	cur = head2
	for {
		if cur == nil {
			return nil
		}
		if findRes(cur, nodeMap) {
			return cur
		} else {
			cur = cur.Next
		}
	}
}

func noLoop(head1, head2 *Node) *Node {
	if head1 == nil || head2 == nil {
		return nil
	}

	n := 0
	cur1 := head1
	cur2 := head2
	for cur1.Next != nil {
		n++
		cur1 = cur1.Next
	}

	for cur2.Next != nil {
		n--
		cur2 = cur2.Next
	}

	if cur1 != cur2 {
		return nil
	}

	if n > 0 {
		cur1 = head1
		cur2 = head2
	} else {
		cur1 = head2
		cur2 = head1
	}

	n = int(math.Abs(float64(n)))

	for n != 0 {
		n--
		cur1 = cur1.Next
	}

	for cur1 != cur2 {
		cur1 = cur1.Next
		cur2 = cur2.Next
	}

	return cur1
}

func bothLoop(head1, loop1, head2, loop2 *Node) *Node {
	if loop1 == loop2 {
		cur1 := head1
		cur2 := head2
		n := 0

		for cur1.Next != loop1 {
			n++
			cur1 = cur1.Next
		}

		for cur2.Next != loop2 {
			n--
			cur2 = cur2.Next
		}

		if n > 0 {
			cur1 = head1
			cur2 = head2
		} else {
			cur1 = head2
			cur2 = head1
		}

		n = int(math.Abs(float64(n)))

		for n != 0 {
			n--
			cur1 = cur1.Next
		}

		for cur1 != cur2 {
			cur1 = cur1.Next
			cur2 = cur2.Next
		}
		return cur1
	} else {
		cur := loop1.Next
		for cur != loop1 {
			if cur == loop2 {
				return loop1
			}
			cur = cur.Next
		}
		return nil
	}
}

func getIntersectNode(head1, head2 *Node) *Node {
	if  head1 == nil || head2 == nil {
		return nil
	}

	loop1 := getLoopNode(head1)
	loop2 := getLoopNode(head2)

	if loop1 == nil && loop2 == nil {
		return noLoop(head1, head2)
	}

	if loop1 != nil && loop2 != nil {
		return bothLoop(head1, loop1, head2, loop2)
	}
	return nil
}

func main() {
	// 1->2->3->4->5->6->7->nil
	head1 := new(Node)
	head1.Val = 1
	head1.Next = new(Node)

	head1.Next.Val = 2
	head1.Next.Next = new(Node)

	head1.Next.Next.Val = 3
	head1.Next.Next.Next = new(Node)

	head1.Next.Next.Next.Val = 4
	head1.Next.Next.Next.Next = new(Node)

	head1.Next.Next.Next.Next.Val = 5
	head1.Next.Next.Next.Next.Next = new(Node)

	head1.Next.Next.Next.Next.Next.Val = 6
	head1.Next.Next.Next.Next.Next.Next = new(Node)

	head1.Next.Next.Next.Next.Next.Next.Val = 7
	head1.Next.Next.Next.Next.Next.Next.Next = nil

	// 0->9->8->6->7->null
	head2 := new(Node)
	head2.Val = 0
	head2.Next = new(Node)

	head2.Next.Val = 9
	head2.Next.Next = new(Node)

	head2.Next.Next.Val = 8
	head2.Next.Next.Next = head1.Next.Next.Next  // 8->4

	fmt.Println(getIntersectNode(head1, head2).Val)


	// 1->2->3->4->5->6->7->4...
	head1 = new(Node)
	head1.Val = 1
	head1.Next = new(Node)

	head1.Next.Val = 2
	head1.Next.Next = new(Node)

	head1.Next.Next.Val = 3
	head1.Next.Next.Next = new(Node)

	head1.Next.Next.Next.Val = 4
	head1.Next.Next.Next.Next = new(Node)

	head1.Next.Next.Next.Next.Val = 5
	head1.Next.Next.Next.Next.Next = new(Node)

	head1.Next.Next.Next.Next.Next.Val = 6
	head1.Next.Next.Next.Next.Next.Next = new(Node)

	head1.Next.Next.Next.Next.Next.Next.Val = 7
	head1.Next.Next.Next.Next.Next.Next.Next = head1.Next.Next.Next

	// 0->9->8->2...
	head2 = new(Node)
	head2.Val = 0
	head2.Next = new(Node)

	head2.Next.Val = 9
	head2.Next.Next = new(Node)

	head2.Next.Next.Val = 8
	head2.Next.Next.Next = head1.Next  // 8->2

	fmt.Println(getIntersectNode(head1, head2).Val)

	// 0->9->8->6->4->5->6...
	head2 = new(Node)
	head2.Val = 0
	head2.Next = new(Node)

	head2.Next.Val = 9
	head2.Next.Next = new(Node)

	head2.Next.Next.Val = 8
	head2.Next.Next.Next = head1.Next.Next.Next.Next.Next  // 8->6

	fmt.Println(getIntersectNode(head1, head2).Val)
}
