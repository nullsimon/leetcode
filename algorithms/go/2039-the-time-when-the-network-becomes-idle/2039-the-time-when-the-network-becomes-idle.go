package thetimewhenthenetworkbecomesidle

import "fmt"

type Node struct {
	ID       int
	Children []*Node
}

func networkBecomesIdle(edges [][]int, patience []int) int {
	seconds := 0
	root := new(Node)
	root.ID = 0
	visted := make(map[int]bool)

	ch := make(chan *Node, 100000)
	ch <- root
	var depth = 1
	for len(ch) > 0 {
		var size = len(ch)
		for j := 0; j < size; j++ {
			cur := <-ch
			for i := 0; i < len(edges); i++ {
				edge := edges[i]
				if edge[0] == cur.ID {
					//avoid circle, but if circle more closed ?
					_, ok := visted[edge[1]]
					if ok {
						continue
					} else {
						visted[edge[1]] = true
					}
					node := new(Node)
					node.ID = edge[1]
					fmt.Println(depth)
					second := depth * 2
					rest := (second - 1) / patience[edge[1]] * patience[edge[1]]
					second += rest

					if second > seconds {
						seconds = second
					}
					ch <- node
					cur.Children = append(cur.Children, node)
				}
			}
		}
		depth++
	}

	return seconds + 1
}
