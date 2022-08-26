package graph

//DFS
func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make([][]int, numCourses)
	for _, pair := range prerequisites {
		graph[pair[0]] = append(graph[pair[0]], pair[1])
	}

	res := make([]int, 0)
	visited, path, hasCycle := make([]bool, numCourses), make([]bool, numCourses), false
	var traverse func(node int)
	traverse = func(node int) {
		if path[node] {
			hasCycle = true
		}

		if visited[node] || hasCycle {
			return
		}

		visited[node] = true
		path[node] = true
		for _, v := range graph[node] {
			traverse(v)
		}
		res = append(res, node)
		path[node] = false
	}

	for i := 0; i < numCourses; i++ {
		traverse(i)
	}

	if hasCycle {
		return []int{}
	}
	return res
}
