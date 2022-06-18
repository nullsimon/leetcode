package deletenodeinalinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	nextNode := node.Next
	node.Next = nextNode.Next
	node.Val = nextNode.Val
}
