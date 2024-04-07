package maximumdepthofnarytree

import "testing"

func TestMaxDepth(t *testing.T) {
	tree := prepareTree()
	got := maxDepth(tree)
	want := 3
	if got != want {
		t.Fatalf(`maxDepth func failed: want %d, got %d`, want, got)
	}
}

func prepareTree() *Node {
	/**
	 *      1
	 *    /  \
	 *   2    3
	 *  / \
	 * 4   5
	 */

	_5 := &Node{Val: 5, Children: []*Node{}}
	_4 := &Node{Val: 4, Children: []*Node{}}
	_3 := &Node{Val: 3, Children: []*Node{}}
	_2 := &Node{Val: 2, Children: []*Node{_4, _5}}
	_1 := &Node{Val: 1, Children: []*Node{_2, _3}}
	return _1
}
