package binarytreelevelordertraversal

import (
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	input := []*TreeNode{prepareTree1()}
	want := [][][]int{{{1}, {2, 3}, {4, 5}}}
	for i := 0; i < len(input); i++ {
		got := levelOrder(input[i])
		if !reflect.DeepEqual(got, want[i]) {
			t.Fatalf(`levelOrder func failed: want %v, got %v`, want[i], got)
		}
	}
}

func prepareTree1() *TreeNode {
	/**
	 *      1
	 *    /  \
	 *   2    3
	 *  / \
	 * 4   5
	 */

	_5 := &TreeNode{Val: 5, Left: nil, Right: nil}
	_4 := &TreeNode{Val: 4, Left: nil, Right: nil}
	_3 := &TreeNode{Val: 3, Left: nil, Right: nil}
	_2 := &TreeNode{Val: 2, Left: _4, Right: _5}
	_1 := &TreeNode{Val: 1, Left: _2, Right: _3}
	return _1
}

func compare(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
