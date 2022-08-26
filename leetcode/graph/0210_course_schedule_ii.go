package graph

//DFS
//func findOrder(numCourses int, prerequisites [][]int) []int {
//	graph := make([][]int, numCourses)
//	for _, pair := range prerequisites {
//		graph[pair[0]] = append(graph[pair[0]], pair[1])
//	}
//
//	res := make([]int, 0)
//	visited, path, hasCycle := make([]bool, numCourses), make([]bool, numCourses), false
//	var traverse func(node int)
//	traverse = func(node int) {
//		if path[node] {
//			hasCycle = true
//		}
//
//		if visited[node] || hasCycle {
//			return
//		}
//
//		visited[node] = true
//		path[node] = true
//		for _, v := range graph[node] {
//			traverse(v)
//		}
//		res = append(res, node)
//		path[node] = false
//	}
//
//	for i := 0; i < numCourses; i++ {
//		traverse(i)
//	}
//
//	if hasCycle {
//		return []int{}
//	}
//	return res
//}

//BFS
func findOrder(numCourses int, prerequisites [][]int) []int {
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

	res := make([]int, 0)
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			q := queue[0]
			res = append(res, q)
			for _, v := range graph[q] {
				indegree[v]--
				if indegree[v] == 0 {
					queue = append(queue, v)
				}
			}
			queue = queue[1:]
		}
	}

	if len(res) == numCourses {
		return res
	}
	return nil
}
