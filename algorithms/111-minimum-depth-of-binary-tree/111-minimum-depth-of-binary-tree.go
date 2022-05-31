package minimumdepthofbinarytree

/*
* sequence or map, or tree. so just using a hard map, if neccesarry
* binary, just split, big to small, one in two, left or right ,big or small, good to know
 */

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

//minDepthBFS using labuladong bfs version code, in go we using a slice
// first add node, then add sub node to list
func minDepthBFS(root *TreeNode) int {
	//base
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	//init a tree queue slice, store in list
	var treeQueue [](*TreeNode)
	var depth = 1
	treeQueue = append(treeQueue, root)
	for treeQueue != nil {
		//use another queue to implicit or just slice remove
		//size := len(treeQueue)
		curQueue := treeQueue
		treeQueue = nil
		for i := 0; i < len(curQueue); i++ {
			cur := curQueue[i]
			//treeQueue = treeQueue[1:]
			if cur.Left == nil && cur.Right == nil {
				return depth
			}
			if cur.Left != nil {
				treeQueue = append(treeQueue, cur.Left)
			}
			if cur.Right != nil {
				treeQueue = append(treeQueue, cur.Right)
			}
		}
		depth++

	}
	return depth

}
