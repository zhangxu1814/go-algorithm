package algorithms4th

import "testing"

func Test_timeTrial(t *testing.T) {
	//t.Log("using time: ", timeTrial(500))

	for N := 250; N <= 8000; N += N {
		time := timeTrial(N)
		t.Logf("%7d, %11d", N, time)
	}
}

func TestDoublingRatio(t *testing.T) {
	prev := timeTrial(125)

	for N := 250; ; N += N {
		time := timeTrial(N)
		t.Logf("%7d, %11d, %5.2f", N, time, float64(time)/float64(prev))
		prev = time
	}
}
