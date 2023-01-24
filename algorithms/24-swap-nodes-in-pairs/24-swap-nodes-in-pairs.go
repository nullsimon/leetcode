package swapnodesinpairs

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummy := new(ListNode)
	tail := dummy
	//odd or single
	if head.Next != nil {
		dummy = head.Next
		tail = head
	} else {
		dummy = head
	}
	p := head
	for p != nil && p.Next != nil {
		n := p.Next
		p.Next = n.Next
		n.Next = p

		if p.Next != nil {
			p = p.Next
		}

		tail.Next = p
		fmt.Println(n, p, tail)
	}

	return dummy
}
