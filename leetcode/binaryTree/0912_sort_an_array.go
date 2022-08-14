package binaryTree

//归并排序
func sortArray(nums []int) []int {
	tmp := make([]int, len(nums))
	var l, r int
	var sort func(int, int)
	var merge func(int, int, int)
	sort = func(lo, hi int) {
		if lo == hi {
			return
		}
		mid := lo + (hi-lo)/2
		sort(lo, mid)
		sort(mid+1, hi)
		merge(lo, mid, hi)
	}
	merge = func(lo, mid, hi int) {
		for i := lo; i <= hi; i++ {
			tmp[i] = nums[i]
		}
		l, r = lo, mid+1
		for i := lo; i <= hi; i++ {
			if l > mid {
				nums[i] = tmp[r]
				r++
			} else if r > hi {
				nums[i] = tmp[l]
				l++
			} else if tmp[l] > tmp[r] {
				nums[i] = tmp[r]
				r++
			} else {
				nums[i] = tmp[l]
				l++
			}
		}
	}

	sort(0, len(nums)-1)
	return nums
}
