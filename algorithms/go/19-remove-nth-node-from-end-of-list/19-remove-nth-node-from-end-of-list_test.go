package removenthnodefromendoflist

import (
	"reflect"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	input := [][]int{
		{1},
		{1, 2, 3, 4, 5},
	}
	nList := []int{
		1,
		2,
	}
	expect := [][]int{
		{},
		{1, 2, 3, 5},
	}
	for i := range input {
		got := removeNthFromEnd(createList(input[i]), nList[i])
		if !reflect.DeepEqual(got, createList(expect[i])) {
			t.Errorf("removeNthFromEnd(%v, %v) = %v; want %v", input[i], nList[i], got, expect[i])
		}
	}
}

func createList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{
		Val: nums[0],
	}
	tmp := head
	for i := 1; i < len(nums); i++ {
		tmp.Next = &ListNode{
			Val: nums[i],
		}
		tmp = tmp.Next
	}
	return head
}
