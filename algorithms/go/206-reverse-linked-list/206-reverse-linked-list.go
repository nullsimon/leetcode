package reverselinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	tail := traverse(head)

	head.Next = nil

	return tail
}

func traverse(node *ListNode) *ListNode {
	if node.Next == nil {
		return node
	}

	tail := traverse(node.Next)
	nextNode := node.Next
	nextNode.Next = node

	return tail
}
