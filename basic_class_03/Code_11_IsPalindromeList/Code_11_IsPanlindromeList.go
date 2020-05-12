package main


import (
	"fmt"
	"github.com/emirpasic/gods/stacks/arraystack"
)

type Node struct {
	data int
	next *Node
}

// need n space
func (node *Node) isPalindromeList1() bool {
	stack := arraystack.New()
	cur := node
	for cur != nil {
		stack.Push(cur.data)
		cur = cur.next
	}

	for node != nil {
		pop, _ := stack.Pop()
		if node.data != pop.(int) {
			return false
		}
		node = node.next
	}
	return true
}

// need n/2 space
func (node *Node) isPalindromeList2() bool {
	right := node.next
	cur := node

	for cur.next != nil && cur.next.next != nil {
		right = right.next
		cur = cur.next.next
	}

	stack := arraystack.New()
	for right != nil {
		stack.Push(right.data)
		right = right.next
	}

	for !stack.Empty() {
		value, _ := stack.Pop()
		if node.data != value.(int) {
			return false
		}
		node = node.next
	}
	return true
}

func (node *Node) isPanlindromeList3() bool {
	fast := node
	slow := node
	for fast.next != nil && fast.next.next != nil {
		fast = fast.next.next
		slow = slow.next
	}

	right := slow.next   // 右边部分的第一个节点
	slow.next = nil     // 将mid节点指向nil，构成左边部分

	pre := new(Node)
	nextNode := new(Node)
	pre = nil
	nextNode = nil
	for right != nil {
		nextNode = right.next
		right.next = pre
		pre = right
		right = nextNode
	}

	last := pre   // 保存最后一个节点
	left := node
	res := true
	for left != nil && pre != nil {
		if left.data != pre.data {
			res = false
			break
		}
		left = left.next
		pre = pre.next
	}

	// 恢复链表
	pre = nil
	//nextNode = nil
	for last != nil {
		nextNode = last.next
		last.next = pre
		pre = last
		last = nextNode
	}
	slow.next = pre
	return res
}

func (node *Node) printSingleList() {
	for node != nil {
		fmt.Printf("%v ",node.data)
		node = node.next
	}
	fmt.Println()
}


/*
func isPalindrome(head *ListNode) bool {
    n := 0
    for p:=head; p!=nil; p=p.Next {
        n ++
    }
    // 找到链表中点的下一个位置
    p := head
    for i:=0; i<n/2; i++ {
        p = p.Next
    }
    if n%2 == 1 {
        p = p.Next
    }
    // 翻转后半段链表
    p = reverse(p)
    for p != nil {
        if head.Val != p.Val {
            return false
        }
        p = p.Next
        head = head.Next
    }
    return true
}
// 反转一个链表，并返回反转后到头节点
func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		head.Next, prev, head = prev, head, head.Next
	}
	return prev
}
*/

func main() {
	head := new(Node)
	head.data = 1

	head.printSingleList()
	fmt.Println(head.isPalindromeList1())
	fmt.Println(head.isPalindromeList2())
	fmt.Println(head.isPanlindromeList3())

	fmt.Println("=====================================")

	head1 := new(Node)
	head1.data = 1

	head1.next = new(Node)
	head1.next.data = 2

	head1.printSingleList()
	fmt.Println(head1.isPalindromeList1())
	fmt.Println(head1.isPalindromeList2())
	fmt.Println(head1.isPanlindromeList3())


	fmt.Println("=====================================")

	head2 := new(Node)
	head2.data = 1

	head2.next = new(Node)
	head2.next.data = 2

	head2.next.next = new(Node)
	head2.next.next.data = 1


	head2.printSingleList()
	fmt.Println(head2.isPalindromeList1())
	fmt.Println(head2.isPalindromeList2())
	fmt.Println(head2.isPanlindromeList3())
	head2.printSingleList()

	fmt.Println("=====================================")

	head3 := new(Node)
	head3.data = 1

	head3.next = new(Node)
	head3.next.data = 2

	head3.next.next = new(Node)
	head3.next.next.data = 3

	head3.next.next.next = new(Node)
	head3.next.next.next.data = 3

	head3.next.next.next.next = new(Node)

	head3.next.next.next.next.data = 2

	head3.next.next.next.next.next = new(Node)
	head3.next.next.next.next.next.data = 1

	head3.printSingleList()
	fmt.Println(head3.isPalindromeList1())
	fmt.Println(head3.isPalindromeList2())
	fmt.Println(head3.isPanlindromeList3())
	head3.printSingleList()



}

