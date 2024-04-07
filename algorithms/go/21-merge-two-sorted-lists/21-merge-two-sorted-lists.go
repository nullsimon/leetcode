package mergetwosortedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	var list = new(ListNode)
	l := list
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			node := new(ListNode)
			node.Val = list1.Val
			if list == nil {
				list = node
			} else {
				list.Next = node
				list = list.Next
			}
			list1 = list1.Next
			continue
		}
		if list1.Val > list2.Val {
			node := new(ListNode)
			node.Val = list2.Val
			if list == nil {
				list = node
			} else {
				list.Next = node
				list = list.Next
			}
			list2 = list2.Next
		}
	}
	if list1 == nil && list2 != nil {
		list.Next = list2
	}
	if list2 == nil && list1 != nil {
		list.Next = list1
	}

	return l.Next
}
