package pathsumii

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		if root.Val == targetSum {
			return [][]int{{root.Val}}
		}
		return nil
	}

	var ch = make(chan *TreeNode, 5000)
	var distCh = make(chan int, 5000)
	var distSlice = make(chan []int, 5000)
	ch <- root
	distCh <- 0
	distSlice <- []int{}

	var res [][]int

	for len(ch) > 0 {
		size := len(ch)
		for i := 0; i < size; i++ {
			cur := <-ch
			parentVal := <-distCh
			//this is a reference slice
			parentSlice := <-distSlice

			curSum := parentVal + cur.Val
			curSlice := append(parentSlice, cur.Val)

			var resSlice = make([]int, len(curSlice))
			copy(resSlice, curSlice)

			if cur.Left == nil && cur.Right == nil {
				if curSum == targetSum {
					res = append(res, resSlice)
				}

			}
			if cur.Left != nil {
				ch <- cur.Left
				distCh <- curSum
				distSlice <- curSlice
			}
			if cur.Right != nil {
				ch <- cur.Right
				distCh <- curSum
				distSlice <- curSlice
			}
		}
	}
	return res
}
