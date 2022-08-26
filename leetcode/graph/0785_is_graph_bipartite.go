package graph

func isBipartite(graph [][]int) bool {
	n := len(graph)
	visited, color, ok := make([]bool, n), make([]bool, n), true

	var traverse func(node int)
	traverse = func(node int) {
		if !ok {
			return
		}

		visited[node] = true
		for _, v := range graph[node] {
			if !visited[v] {
				color[v] = !color[node]
				traverse(v)
			} else {
				if color[node] == color[v] {
					ok = false
					return
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		if !visited[i] {
			traverse(i)
		}
	}

	return ok
}
