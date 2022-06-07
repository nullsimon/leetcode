package pathsumii

import (
	"reflect"
	"testing"
)

func TestPathSum(t *testing.T) {
	root := prepareTree1()
	target := 22
	want := [][]int{{5, 8, 4, 5}}
	got := pathSum(root, target)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf(`hasPathSum function failed: want %v, got %v`, want, got)
	}
}

func prepareTree() *TreeNode {
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

func prepareTree1() *TreeNode {
	_7 := &TreeNode{Val: 5, Left: nil, Right: nil}
	_6 := &TreeNode{Val: 1, Left: nil, Right: nil}
	_5 := &TreeNode{Val: 4, Left: _7, Right: _6}
	_4 := &TreeNode{Val: 13, Left: nil, Right: nil}
	_3 := &TreeNode{Val: 8, Left: _4, Right: _5}
	_2 := &TreeNode{Val: 5, Left: nil, Right: _3}
	return _2
}
