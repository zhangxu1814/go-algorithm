package graph

//DFS
//func isBipartite(graph [][]int) bool {
//	n := len(graph)
//	visited, color, ok := make([]bool, n), make([]bool, n), true
//
//	var traverse func(node int)
//	traverse = func(node int) {
//		if !ok {
//			return
//		}
//
//		visited[node] = true
//		for _, v := range graph[node] {
//			if !visited[v] {
//				color[v] = !color[node]
//				traverse(v)
//			} else {
//				if color[node] == color[v] {
//					ok = false
//					return
//				}
//			}
//		}
//	}
//
//	for i := 0; i < n; i++ {
//		if !visited[i] {
//			traverse(i)
//		}
//	}
//
//	return ok
//}

//BFS
func isBipartite(graph [][]int) bool {
	visited, color, ok := make([]bool, len(graph)), make([]bool, len(graph)), true

	var bfs func(node int)
	bfs = func(node int) {
		queue := []int{node}
		visited[node] = true

		for len(queue) > 0 {
			q := queue[0]
			for _, v := range graph[q] {
				if !visited[v] {
					color[v] = !color[q]
					visited[v] = true
					queue = append(queue, v)
				} else {
					if color[v] == color[q] {
						ok = false
						return
					}
				}
			}
			queue = queue[1:]
		}
	}

	for i := 0; i < len(graph); i++ {
		if ok && !visited[i] {
			bfs(i)
		}
	}

	return ok
}
