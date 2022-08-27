package graph

//DFS
func possibleBipartition(n int, dislikes [][]int) bool {
	n++
	graph := make([][]int, n)
	for _, pair := range dislikes {
		graph[pair[0]] = append(graph[pair[0]], pair[1])
		graph[pair[1]] = append(graph[pair[1]], pair[0])
	}

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
				if color[v] == color[node] {
					ok = false
					return
				}
			}
		}
	}

	for i := 1; i < n; i++ {
		if ok && !visited[i] {
			traverse(i)
		}
	}

	return ok
}
