package maximumdepthofbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	var left, right = 0, 0
	if root.Left != nil {
		left = maxDepth(root.Left)
	}
	if root.Right != nil {
		right = maxDepth(root.Right)
	}
	var res = 0
	if left > right {
		res = left
	} else {
		res = right
	}
	return res + 1
}
