
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, L int, R int) int {
    sum := 0
     if root != nil {
        if root.Val >= L && root.Val <= R {
            sum += root.Val
        }
        return sum + rangeSumBST(root.Left, L, R) + rangeSumBST(root.Right, L, R)

    }
    return sum
}

