package binarytreemaximumpathsum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}

	var left, right = 0, 0
	var res = root.Val

	if root.Left != nil {
		left = maxPathSum(root.Left)
	}
	if root.Right != nil {
		right = maxPathSum(root.Right)
	}
	if root.Left == nil {
		return maxNodeTwo(right, res)
	}
	if root.Right == nil {
		return maxNodeTwo(left, res)

	}

	res = maxNode(left, right, res)
	return res
}

func maxNodeTwo(a, c int) int {
	if a < 0 && c < 0 {
		if a > c {
			return a
		}
		return c
	}
	if a > a+c {
		return a
	}
	if c > a+c {
		return c
	}
	return a + c
}

func maxNode(a, b, c int) int {

	if a > a+b+c {
		return a
	}
	if b > a+b+c {
		return b
	}
	if c > a+b+c {
		return c
	}
	if a+c > a+b+c && a+c > c {
		return a + c
	}
	if b+c > a+b+c && b+c > c {
		return b + c
	}

	return a + b + c
}
