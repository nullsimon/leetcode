package narytreepreordertraversal

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	if len(root.Children) == 0 {
		return [][]int{{root.Val}}
	}
	var ch = make(chan *Node, 10000)
	ch <- root
	var leverSlice [][]int
	for len(ch) > 0 {
		var size = len(ch)
		var curSlice []int
		for i := 0; i < size; i++ {
			cur := <-ch
			curSlice = append(curSlice, cur.Val)

			if len(cur.Children) == 0 {
				continue
			} else if len(cur.Children) != 0 {
				for j := 0; j < len(cur.Children); j++ {
					ch <- cur.Children[j]
				}
			}

		}
		leverSlice = append(leverSlice, curSlice)
	}
	return leverSlice
}
