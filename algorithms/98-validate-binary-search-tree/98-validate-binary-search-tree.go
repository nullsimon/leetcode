package validatebinarysearchtree

import (
	"reflect"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	s := traverse(root)
	var target []int
	target = append(target, s...)
	sort.Ints(s)
	if !reflect.DeepEqual(s, target) {
		return false
	}
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return false
		}
	}
	return true
}

func isValidBST2(root *TreeNode, min *TreeNode, max *TreeNode) bool {
	if root == nil {
		return true
	}
	if min != nil && root.Val < 
}

func traverse(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	left := traverse(root.Left)
	right := traverse(root.Right)
	var res []int
	if len(left) > 0 {
		res = append(res, left...)
	}
	res = append(res, root.Val)
	if len(right) > 0 {
		res = append(res, right...)

	}
	return res
}
