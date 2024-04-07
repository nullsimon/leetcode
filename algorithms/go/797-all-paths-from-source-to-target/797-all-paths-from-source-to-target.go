package allpathsfromsourcetotarget

import "fmt"

func allPathsSourceTarget(graph [][]int) [][]int {
	var res [][]int
	var path []int
	traverse(graph, 0, path, &res)
	return res
}

func traverse(graph [][]int, s int, path []int, res *([][]int)) {

	path = append(path, s)
	resPath := make([]int, len(path))
	copy(resPath, path)

	n := len(graph)
	if s == n-1 {
		//end
		*res = append(*res, resPath)
		fmt.Println(resPath)
	}
	for i := 0; i < len(graph[s]); i++ {
		traverse(graph, graph[s][i], resPath, res)
	}
}

func allPathsSourceTarget_BFS(graph [][]int) [][]int {
	if graph == nil {
		return nil
	}
	var ch = make(chan int, 225)
	var pathCh = make(chan []int, 225)
	var res [][]int
	ch <- 0
	pathCh <- []int{}
	for len(ch) > 0 {
		for i := 0; i < len(ch); i++ {
			cur := <-ch
			parents := <-pathCh

			curPath := append(parents, cur)
			var resPath = make([]int, len(curPath))
			copy(resPath, curPath)
			if len(graph[cur]) == 0 {
				res = append(res, resPath)
			}
			for j := 0; j < len(graph[cur]); j++ {
				ch <- graph[cur][j]
				pathCh <- resPath
			}
		}
	}
	return res
}
