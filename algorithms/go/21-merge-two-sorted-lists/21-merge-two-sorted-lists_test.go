package mergetwosortedlists

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	list1 := prepareList1()
	list2 := prepareList2()
	want := prepareList3()
	got := mergeTwoLists(list1, list2)
	if !reflect.DeepEqual(want, got) {
		t.Fatalf(`merge function not working as expected, want %v, got %v`, want, got)
	}
}

func prepareList1() *ListNode {
	l3 := &ListNode{
		Val: 4,
	}
	l2 := &ListNode{
		Val:  2,
		Next: l3,
	}
	l1 := &ListNode{
		Val:  1,
		Next: l2,
	}
	return l1
}

func prepareList2() *ListNode {
	l3 := &ListNode{
		Val: 4,
	}
	l2 := &ListNode{
		Val:  3,
		Next: l3,
	}
	l1 := &ListNode{
		Val:  1,
		Next: l2,
	}
	return l1
}

func prepareList3() *ListNode {
	l6 := &ListNode{
		Val: 4,
	}
	l5 := &ListNode{
		Val:  4,
		Next: l6,
	}
	l4 := &ListNode{
		Val:  3,
		Next: l5,
	}
	l3 := &ListNode{
		Val:  2,
		Next: l4,
	}
	l2 := &ListNode{
		Val:  1,
		Next: l3,
	}
	l1 := &ListNode{
		Val:  1,
		Next: l2,
	}
	return l1
}

func compare(l1 *ListNode, l2 *ListNode) bool {
	if l1 == nil && l2 == nil {
		return true
	}
	if l1 == nil || l2 == nil {
		return false
	}
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	if l1 == nil && l2 == nil {
		return true
	}
	return false
}
