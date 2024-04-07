package preordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	var ints []int
	ints = append(ints, root.Val)
	if root.Left != nil {
		ints = append(ints, preorderTraversal(root.Left)...)
	}
	if root.Right != nil {
		ints = append(ints, preorderTraversal(root.Right)...)
	}
	return ints
}
