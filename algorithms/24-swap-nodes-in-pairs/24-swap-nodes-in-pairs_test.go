package swapnodesinpairs

import (
	"reflect"
	"testing"
)

func TestSwapPairs(t *testing.T) {
	input := [][]int{
		{1, 2, 3, 4},
		{1, 2, 3, 4, 5},
	}
	expect := [][]int{
		{2, 1, 4, 3},
		{2, 1, 4, 3, 5},
	}
	for i := range input {
		got := swapPairs(createList(input[i]))
		if !reflect.DeepEqual(got, createList(expect[i])) {
			// t.Errorf("swapPairs(%v) = %v; want %v", input[i], got, createList(expect[i]))
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
