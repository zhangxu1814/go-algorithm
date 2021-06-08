package algorithms4th

func count(a []int) int {
	N := len(a)
	cnt := 0

	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			for k := j + 1; k < N; k++ {
				if a[i]+a[j]+a[k] == 0 {
					cnt++
				}
			}
		}
	}

	return cnt
}
