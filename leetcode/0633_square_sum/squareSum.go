package square_sum

import "math"

func judgeSquareSum(c int) bool {
	count := int(math.Sqrt(float64(c)))
	for i := 0; i <= count; i++ {
		if res := math.Sqrt(float64(c - i*i)); res-float64(int(res)) == 0 {
			return true
		}
	}

	return false
}
