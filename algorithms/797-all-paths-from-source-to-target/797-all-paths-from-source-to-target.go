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
