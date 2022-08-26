package graph

//DFS
//func canFinish(numCourses int, prerequisites [][]int) bool {
//	graph := make([][]int, numCourses)
//	for _, pair := range prerequisites {
//		graph[pair[1]] = append(graph[pair[1]], pair[0])
//	}
//
//	hasCycle := false
//	visited, path := make([]bool, numCourses), make([]bool, numCourses)
//
//	var traverse func(node int)
//	traverse = func(node int) {
//		if path[node] {
//			hasCycle = true
//		}
//		if visited[node] || hasCycle {
//			return
//		}
//
//		visited[node] = true
//		path[node] = true
//		for _, v := range graph[node] {
//			traverse(v)
//		}
//		path[node] = false
//	}
//
//	for i := 0; i < numCourses; i++ {
//		traverse(i)
//	}
//
//	return !hasCycle
//}

//BFS
func canFinish(numCourses int, prerequisites [][]int) bool {
	graph, indegree := make([][]int, numCourses), make([]int, numCourses)
	for _, pair := range prerequisites {
		graph[pair[1]] = append(graph[pair[1]], pair[0])
		indegree[pair[0]]++
	}

	queue := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	count := 0
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			q := queue[0]
			count++
			for _, v := range graph[q] {
				indegree[v]--
				if indegree[v] == 0 {
					queue = append(queue, v)
				}
			}
			queue = queue[1:]
		}
	}

	if count == numCourses {
		return true
	}
	return false
}
