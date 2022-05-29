package balancedbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {

	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}

	var left, right = 0, 0
	balance := true

	if root.Left != nil {
		left, balance = walkTree(root.Left)
		if !balance {
			return false
		}
	}
	if root.Right != nil {
		right, balance = walkTree(root.Right)
		if !balance {
			return false
		}
	}
	if left-right > 1 || right-left > 1 {
		return false
	}
	return true
}

func walkTree(t *TreeNode) (int, bool) {

	if t == nil {
		return 0, true
	}
	if t.Left == nil && t.Right == nil {
		return 1, true
	}
	var left, right = 0, 0
	balance := true

	if t.Left != nil {
		left, balance = walkTree(t.Left)
		if !balance {
			return 0, false
		}
	}
	if t.Right != nil {
		right, balance = walkTree(t.Right)
		if !balance {
			return 0, false
		}
	}

	if left-right > 1 || right-left > 1 {
		return 0, false
	}
	return max(left, right) + 1, true
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
