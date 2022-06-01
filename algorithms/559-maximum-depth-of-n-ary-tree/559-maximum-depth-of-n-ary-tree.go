package maximumdepthofnarytree

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	if len(root.Children) == 0 {
		return 1
	}
	var depth = 0
	var ch = make(chan *Node, 10000)
	ch <- root
	for len(ch) > 0 {

		var size = len(ch)
		for i := 0; i < size; i++ {
			cur := <-ch
			if len(cur.Children) == 0 {
				continue
			}
			for j := 0; j < len(cur.Children); j++ {
				ch <- cur.Children[j]
			}
		}
		depth++
	}
	return depth
}
