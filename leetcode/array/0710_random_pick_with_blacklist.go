package array

//import (
//	"math/rand"
//	"time"
//)

//type Solution struct {
//	n        int
//	location map[int]int
//}
//
//func Constructor(n int, blacklist []int) Solution {
//	rand.Seed(time.Now().UnixNano())
//	location := map[int]int{}
//	for _, v := range blacklist {
//		location[v] = v
//	}
//	cur, l := n-1, n-len(blacklist)-1
//	for _, v := range blacklist {
//		if v > l {
//			continue
//		}
//		for location[cur] == cur {
//			cur--
//		}
//		location[v] = cur
//		cur--
//	}
//
//	return Solution{
//		n:        n - len(blacklist),
//		location: location,
//	}
//}
//
//func (this *Solution) Pick() int {
//	r := rand.Intn(this.n)
//	if v, exi := this.location[r]; exi {
//		return v
//	} else {
//		return r
//	}
//}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(n, blacklist);
 * param_1 := obj.Pick();
 */
