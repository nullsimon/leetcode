package addtwonumbers

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	//version1 using int conversion, maybe failed to big int
	// l1Int := toNumber(l1)
	// l2Int := toNumber(l2)
	// total := l1Int + l2Int
	// return toListNode(total)

	// just use simple overFlow detection and reverse List algo

	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil && l2 != nil {
		return l2
	}
	if l1 != nil && l2 == nil {
		return l1
	}
	var l3 *ListNode
	overFlow := 0
	for l1 != nil || l2 != nil {
		l := new(ListNode)
		val := 0
		if l1 != nil {
			val += l1.Val
		}
		if l2 != nil {
			val += l2.Val
		}
		val += overFlow
		if val/10 >= 1 {
			overFlow = 1
		} else {
			overFlow = 0
		}
		l.Val = val % 10
		if l3 == nil {
			l.Next = nil
		} else {
			l.Next = l3
		}
		l3 = l
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if overFlow > 0 {
		l := new(ListNode)
		l.Val = overFlow
		l.Next = l3
		l3 = l
	}
	return reverse(l3)

}

func reverse(l *ListNode) *ListNode {
	if l == nil {
		return nil
	}
	var nl *ListNode
	for l != nil {
		ll := new(ListNode)
		ll.Val = l.Val
		if nl == nil {
			ll.Next = nil
		} else {
			ll.Next = nl
		}
		nl = ll
		l = l.Next
	}
	return nl
}

//toNumber wail be failed if the number beyond 2^64 limit,may be to slice and add them up
func toNumber(l *ListNode) int {
	var nums int
	var base = 1
	for l != nil {
		nums = nums + l.Val*base
		base = base * 10
		l = l.Next
	}
	return nums
}

func toListNode(i int) *ListNode {
	var pre *ListNode
	for i > 0 {
		val := i % 10
		i = i / 10
		nl := new(ListNode)
		nl.Val = val
		if pre == nil {
			nl.Next = nil
		} else {
			nl.Next = pre
		}
		pre = nl
	}
	for pre != nil {
		n := new(ListNode)
		n.Val = pre.Val
		pre = pre.Next
		n.Next = pre
	}

	return pre
}
