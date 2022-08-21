package binaryTree

func findKthLargest(nums []int, k int) int {
	var partition func(lo, hi int) int
	partition = func(lo, hi int) int {
		p, l, r := nums[lo], lo+1, hi

		for l <= r {
			for l < hi && nums[l] <= p {
				l++
			}
			for r > lo && nums[r] > p {
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

	lo, hi, k := 0, len(nums)-1, len(nums)-k
	for lo <= hi {
		p := partition(lo, hi)
		if p == k {
			return nums[k]
		}
		if p > k {
			hi = p - 1
		} else {
			lo = p + 1
		}
	}

	return -1
}
