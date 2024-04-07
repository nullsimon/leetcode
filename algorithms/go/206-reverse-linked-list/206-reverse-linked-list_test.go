package reverselinkedlist

import (
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	input := prepareList1()
	want := prepareList2()
	got := reverseList(input)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf(`reverseList function failed, want %v, got %v`, want, got)
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
		Val: 1,
	}
	l2 := &ListNode{
		Val:  2,
		Next: l3,
	}
	l1 := &ListNode{
		Val:  4,
		Next: l2,
	}
	return l1
}
