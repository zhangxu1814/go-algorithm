package algorithms4th

import (
	"math/rand"
	"time"
)

func timeTrial(N int) int64 {
	var MAX int = 1000000
	a := make([]int, N)

	for i := 0; i < N; i++ {
		a[i] = rand.Intn(MAX*2) - MAX
	}

	t1 := time.Now()
	count(a)
	t2 := time.Now()

	return int64(t2.Sub(t1).Microseconds())
}
