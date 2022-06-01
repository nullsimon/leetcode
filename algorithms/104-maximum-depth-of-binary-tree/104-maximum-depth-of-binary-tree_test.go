package maximumdepthofbinarytree

import "testing"

func TestMaxDepth(t *testing.T) {
	tree := prepareTree()
	got := maxDepth(tree)
	want := 3
	if got != want {
		t.Fatalf(`maxDepth func failed: want %d, got %d`, want, got)
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
