package postordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	var ints []int
	if root.Left != nil {
		ints = append(ints, postorderTraversal(root.Left)...)
	}
	if root.Right != nil {
		ints = append(ints, postorderTraversal(root.Right)...)
	}
	ints = append(ints, root.Val)
	return ints
}
