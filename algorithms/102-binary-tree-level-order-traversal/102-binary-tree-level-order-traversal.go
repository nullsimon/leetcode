package binarytreelevelordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return [][]int{{root.Val}}
	}

	var treeQueue [](*TreeNode)
	treeQueue = append(treeQueue, root)
	var leverTreeSlice [][]int
	for size := len(treeQueue); size > 0; {
		var treeSlice []int
		for i := 0; i < size; i++ {
			//deque enque
			cur := treeQueue[0]
			treeQueue = treeQueue[1:]
			treeSlice = append(treeSlice, cur.Val)
			if cur.Left == nil && cur.Right == nil {
				continue
			}
			if cur.Left != nil {
				treeQueue = append(treeQueue, cur.Left)
			}
			if cur.Right != nil {
				treeQueue = append(treeQueue, cur.Right)
			}
		}
		leverTreeSlice = append(leverTreeSlice, treeSlice)
	}

	return leverTreeSlice
}
