package narytreepreordertraversal

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	if root == nil {
		return nil
	}
	if len(root.Children) == 0 {
		return []int{root.Val}
	}
	var s []int
	s = append(s, root.Val)
	for i := 0; i < len(root.Children); i++ {
		s = append(s, preorder(root.Children[i])...)
	}
	return s
}
