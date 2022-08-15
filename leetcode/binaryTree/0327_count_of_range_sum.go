package binaryTree

func countRangeSum(nums []int, lower int, upper int) int {
	preSum := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}
	nums = preSum

	tmp := make([]int, len(nums))
	var l, r, start, end, count int
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
		start, end = mid+1, mid+1
		for i := lo; i <= mid; i++ {
			for start <= hi && nums[start]-nums[i] < lower {
				start++
			}
			end = start
			for end <= hi && nums[end]-nums[i] <= upper {
				end++
			}
			count += end - start
		}

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
	return count
}
