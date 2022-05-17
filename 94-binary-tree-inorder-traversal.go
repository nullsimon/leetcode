func inorderTraversal(root *TreeNode) []int {

	nums := []int{}
	if root == nil {
		return nums
	}
	if root.Left != nil {
		nums = append(nums, inorderTraversal(root.Left)...)
	}
	nums = append(nums, root.Val)
	if root.Right != nil {
		nums = append(nums, inorderTraversal(root.Right)...)
	}
	return nums
}
