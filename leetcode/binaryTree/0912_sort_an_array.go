package binaryTree

import "math/rand"

//归并排序
//func sortArray(nums []int) []int {
//	tmp := make([]int, len(nums))
//	var l, r int
//	var sort func(int, int)
//	var merge func(int, int, int)
//	sort = func(lo, hi int) {
//		if lo == hi {
//			return
//		}
//		mid := lo + (hi-lo)/2
//		sort(lo, mid)
//		sort(mid+1, hi)
//		merge(lo, mid, hi)
//	}
//	merge = func(lo, mid, hi int) {
//		for i := lo; i <= hi; i++ {
//			tmp[i] = nums[i]
//		}
//		l, r = lo, mid+1
//		for i := lo; i <= hi; i++ {
//			if l > mid {
//				nums[i] = tmp[r]
//				r++
//			} else if r > hi {
//				nums[i] = tmp[l]
//				l++
//			} else if tmp[l] > tmp[r] {
//				nums[i] = tmp[r]
//				r++
//			} else {
//				nums[i] = tmp[l]
//				l++
//			}
//		}
//	}
//
//	sort(0, len(nums)-1)
//	return nums
//}

//快速排序
func sortArray(nums []int) []int {
	var qs func(lo, hi int)
	var partition func(lo, hi int) int
	qs = func(lo, hi int) {
		if lo >= hi {
			return
		}

		p := partition(lo, hi)
		qs(lo, p-1)
		qs(p+1, hi)
	}
	partition = func(lo, hi int) int {
		n, l, r := nums[lo], lo+1, hi

		for l <= r {
			for l < hi && nums[l] <= n {
				l++
			}
			for r > lo && nums[r] > n {
				r--
			}
			if l >= r {
				break
			}
			nums[l], nums[r] = nums[r], nums[l]
		}
		nums[lo], nums[r] = nums[r], nums[lo]
		return r
	}

	n := len(nums)
	for i := 0; i < n; i++ {
		r := i + rand.Intn(n-i)
		nums[i], nums[r] = nums[r], nums[i]
	}
	qs(0, len(nums)-1)

	return nums
}
