package minimumdepthofbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}

	res := 1
	var left, right = 1, 1
	if root.Right != nil && root.Left == nil {
		right = minDepth(root.Right)
		return right + 1
	}

	if root.Left != nil && root.Right == nil {
		left = minDepth(root.Left)
		return left + 1
	}

	right = minDepth(root.Right)
	left = minDepth(root.Left)

	res = min(left, right) + 1

	return res

}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
