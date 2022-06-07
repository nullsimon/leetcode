package pathsum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}
	var ch = make(chan *TreeNode, 5000)
	var distCh = make(chan int, 5000)
	ch <- root
	distCh <- 0
	for len(ch) > 0 {
		size := len(ch)
		for i := 0; i < size; i++ {
			cur := <-ch
			parentVal := <-distCh
			curSum := parentVal + cur.Val

			if cur.Left == nil && cur.Right == nil {
				if curSum == targetSum {
					return true
				}
			}
			if cur.Left != nil {
				ch <- cur.Left
				distCh <- curSum
			}
			if cur.Right != nil {
				ch <- cur.Right
				distCh <- curSum
			}
		}
	}
	return false
}
