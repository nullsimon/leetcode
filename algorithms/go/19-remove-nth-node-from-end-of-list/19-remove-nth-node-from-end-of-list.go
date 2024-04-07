package removenthnodefromendoflist

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// recursion
	if head.Next == nil {
		return nil
	}
	traverse(head, &n)
	if n > 0 {
		head = head.Next
	}
	return head

}
func traverse(head *ListNode, n *int) *ListNode {
	if head == nil {
		return nil
	}
	next := head.Next
	if next != nil {
		traverse(next, n)
		*n--
		if *n == 0 {
			next = head.Next
			head.Next = next.Next
			return nil
		}
	}
	return nil
}
