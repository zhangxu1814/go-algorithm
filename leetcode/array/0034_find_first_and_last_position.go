package array

func searchRange(nums []int, target int) []int {
	return []int{left(nums, target), right(nums, target)}
}

func left(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] >= target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	if l == len(nums) {
		return -1
	}
	if nums[l] == target {
		return l
	} else {
		return -1
	}
}

func right(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] <= target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	if r < 0 {
		return -1
	}
	if nums[r] == target {
		return r
	} else {
		return -1
	}
}
