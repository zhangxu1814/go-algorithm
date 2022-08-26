package graph

func allPathsSourceTarget(graph [][]int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	var traverse func(node int)
	traverse = func(node int) {
		path = append(path, node)
		if node == len(graph)-1 {
			res = append(res, append([]int{}, path...))
		}
		for i := 0; i < len(graph[node]); i++ {
			traverse(graph[node][i])
		}
		path = path[:len(path)-1]
	}

	traverse(0)
	return res
}
