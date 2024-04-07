package validatebinarysearchtree

import "testing"

func TestIsValidBST(t *testing.T) {
	input := []*TreeNode{
		prepareTree1(),
		prepareTree2(),
	}
	want := []bool{
		true,
		false,
	}
	for i := 0; i < len(input); i++ {
		got := isValidBST(input[i])
		if got != want[i] {
			t.Fatalf(`isValidBST function failed, want %v, got %v`, want[i], got)
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

	_3 := &TreeNode{Val: 3, Left: nil, Right: nil}
	_2 := &TreeNode{Val: 1, Left: nil, Right: nil}
	_1 := &TreeNode{Val: 2, Left: _2, Right: _3}
	return _1
}

func prepareTree2() *TreeNode {
	/**
	 *      1
	 *    /  \
	 *   2    3
	 *  / \
	 * 4   5
	 */

	_3 := &TreeNode{Val: 2, Left: nil, Right: nil}
	_2 := &TreeNode{Val: 2, Left: nil, Right: nil}
	_1 := &TreeNode{Val: 2, Left: _2, Right: _3}
	return _1
}
