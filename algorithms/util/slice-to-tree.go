package util

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// contruct a lever traversal slice to treenode
func sliceToTree(s []int) *TreeNode {
	if len(s) == 0 {
		return nil
	}
	var root = new(TreeNode)
	root.Val = s[0]
	if len(s) == 1 {
		return root
	}
	var treeQueue [](*TreeNode)
	treeQueue = append(treeQueue, root)
	base := 2
	for i := 1; i < len(s); i++ {
		for j := 0; j < base && j < len(s); j++ {

		}
		base = base * 2
	}
	return root
}

func ReConstructBinaryTree(pre, in []int) *TreeNode {
	if nil == pre || nil == in {
		return nil
	}

	return reConstructBinaryTree(pre, in, 0, len(pre)-1, 0, len(in)-1)
}
func reConstructBinaryTree(pre, in []int, preStart, preEnd, inStart, inEnd int) *TreeNode {
	tree := &TreeNode{Val: pre[preStart], Left: nil, Right: nil}
	if preStart == preEnd && inStart == inEnd {
		return tree
	}
	rootIdx := inStart
	for rootIdx < inEnd {
		if pre[preStart] == in[rootIdx] {
			break
		}
		rootIdx++
	}
	leftLength, rightLength := rootIdx-inStart, inEnd-rootIdx
	if leftLength > 0 {
		tree.Left = reConstructBinaryTree(pre, in, preStart+1, preStart+leftLength, inStart, rootIdx-1)
	}
	if rightLength > 0 {
		tree.Right = reConstructBinaryTree(pre, in, preStart+leftLength+1, preEnd, rootIdx+1, inEnd)
	}
	return tree
}
