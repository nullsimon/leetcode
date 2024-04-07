package invertbinarytree

import (
	"reflect"
	"testing"
)

func TestInvertTree(t *testing.T) {
	input := prepareTree()
	want := prepareTree1()
	got := invertTree(input)
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
	_2 := &TreeNode{Val: 2, Left: _5, Right: _4}
	_1 := &TreeNode{Val: 1, Left: _3, Right: _2}
	return _1
}
