package deletenodeinalinkedlist

import (
	"reflect"
	"testing"
)

func TestDeleteNode(t *testing.T) {
	input := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 9,
				},
			},
		},
	}
	want := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 9,
			},
		},
	}
	deleteNode(input.Next)
	if !reflect.DeepEqual(want, input) {
		t.Fatalf(`deleteNode function failed, want %v, got %v`, want, input)
	}
}
