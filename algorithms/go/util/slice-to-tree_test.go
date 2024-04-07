package util

import "testing"

func TestSliceToTree(t *testing.T) {
	pre, in := []int{1, 2, 4, 7, 3, 5, 6, 8}, []int{4, 7, 2, 1, 5, 3, 8, 6}
	node := ReConstructBinaryTree(pre, in)

	//got := intToRoman(input[i])
	// got := levelOrder(input[i])
	// if got != want[i] {
	// 	t.Fatalf(`integer to roman func failed: want %s, got %s`, want[i], got)
	// }
	t.Log(`this is what i want`, node)

}
