package removelinkedlistelements

import (
	"reflect"
	"testing"
)

func TestRemoveElements(t *testing.T) {
	input := []*ListNode{
		prepareList1(),
		prepareList3(),
	}
	input2 := []int{
		1, 7,
	}
	want := []*ListNode{
		prepareList2(),
		nil,
	}
	for i := range input {
		got := removeElements(input[i], input2[i])
		if !reflect.DeepEqual(got, want[i]) {
			t.Fatalf(`removeElements function  failed, want %v, got %v .`, want[i], got)
		}
	}

}

func prepareList1() *ListNode {
	l1 := &ListNode{
		Val:  1,
		Next: nil,
	}
	l2 := &ListNode{
		Val:  2,
		Next: l1,
	}
	l3 := &ListNode{
		Val:  3,
		Next: l2,
	}
	l4 := &ListNode{
		Val:  4,
		Next: l3,
	}
	l5 := &ListNode{
		Val:  5,
		Next: l4,
	}
	return l5
}

func prepareList2() *ListNode {

	l2 := &ListNode{
		Val:  2,
		Next: nil,
	}
	l3 := &ListNode{
		Val:  3,
		Next: l2,
	}
	l4 := &ListNode{
		Val:  4,
		Next: l3,
	}
	l5 := &ListNode{
		Val:  5,
		Next: l4,
	}
	return l5
}

func prepareList3() *ListNode {

	l2 := &ListNode{
		Val:  7,
		Next: nil,
	}
	l3 := &ListNode{
		Val:  7,
		Next: l2,
	}
	l4 := &ListNode{
		Val:  7,
		Next: l3,
	}
	l5 := &ListNode{
		Val:  7,
		Next: l4,
	}
	return l5
}
