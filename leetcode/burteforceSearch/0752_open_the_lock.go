package burteforceSearch

//BFS
//func openLock(deadends []string, target string) int {
//	up := func(s string, n int) string {
//		tmp := []byte(s)
//		if tmp[n] == '9' {
//			tmp[n] = '0'
//		} else {
//			tmp[n]++
//		}
//		return string(tmp)
//	}
//	down := func(s string, n int) string {
//		tmp := []byte(s)
//		if tmp[n] == '0' {
//			tmp[n] = '9'
//		} else {
//			tmp[n]--
//		}
//		return string(tmp)
//	}
//
//	visited := make(map[string]bool)
//	for _, s := range deadends {
//		visited[s] = true
//	}
//	if visited["0000"] {
//		return -1
//	}
//	visited["0000"] = true
//
//	queue := []string{"0000"}
//	count := 0
//
//	for len(queue) > 0 {
//		n := len(queue)
//		for i := 0; i < n; i++ {
//			q := queue[0]
//			queue = queue[1:]
//			if q == target {
//				return count
//			}
//			for j := 0; j < 4; j++ {
//				if s := up(q, j); !visited[s] {
//					queue = append(queue, s)
//					visited[s] = true
//				}
//				if s := down(q, j); !visited[s] {
//					queue = append(queue, s)
//					visited[s] = true
//				}
//			}
//		}
//		count++
//	}
//
//	return -1
//}

//双向BFS
func openLock(deadends []string, target string) int {
	up := func(s string, n int) string {
		tmp := []byte(s)
		if tmp[n] == '9' {
			tmp[n] = '0'
		} else {
			tmp[n]++
		}
		return string(tmp)
	}
	down := func(s string, n int) string {
		tmp := []byte(s)
		if tmp[n] == '0' {
			tmp[n] = '9'
		} else {
			tmp[n]--
		}
		return string(tmp)
	}

	visited := make(map[string]bool)
	for _, s := range deadends {
		visited[s] = true
	}
	if visited["0000"] {
		return -1
	}
	q1, q2 := make(map[string]bool), make(map[string]bool)
	q1["0000"] = true
	q2[target] = true
	count := 0

	for len(q1) > 0 && len(q2) > 0 {
		tmp := make(map[string]bool)
		for n := range q1 {
			if q2[n] {
				return count
			}
			visited[n] = true
			for j := 0; j < 4; j++ {
				if s := up(n, j); !visited[s] {
					tmp[s] = true
				}
				if s := down(n, j); !visited[s] {
					tmp[s] = true
				}
			}
		}
		count++
		q1, q2 = q2, tmp
	}

	return -1
}
