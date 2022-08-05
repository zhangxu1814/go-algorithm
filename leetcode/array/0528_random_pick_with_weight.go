package array

import (
	"math/rand"
	"time"
)

type Solution struct {
	preSum []int
}

//func Constructor(w []int) Solution {
//	preSum := make([]int, len(w)+1)
//	preSum[0] = 0
//	for i := 1; i <= len(w); i++ {
//		preSum[i] = preSum[i-1] + w[i-1]
//	}
//	return Solution{preSum: preSum}
//}

func (this *Solution) PickIndex() int {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(this.preSum[len(this.preSum)-1]) + 1
	l, r := 0, len(this.preSum)-1
	for l <= r {
		mid := l + (r-l)/2
		if this.preSum[mid] >= random {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l - 1
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
