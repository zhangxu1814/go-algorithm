package graph

//DFS
func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	for _, pair := range prerequisites {
		graph[pair[1]] = append(graph[pair[1]], pair[0])
	}

	hasCycle := false
	visited, path := make([]bool, numCourses), make([]bool, numCourses)

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
		path[node] = false
	}

	for i := 0; i < numCourses; i++ {
		traverse(i)
	}

	return !hasCycle
}
