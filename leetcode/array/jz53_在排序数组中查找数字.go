package array

//func search(nums []int, target int) int {
//	l, r := findLeft(nums, target), findRight(nums, target)
//	if l == -1 {
//		return 0
//	}
//	return r - l + 1
//}

func findLeft(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)/2
		if nums[m] == target {
			r = m - 1
		} else if nums[m] > target {
			r = m - 1
		} else {
			l = m + 1
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

func findRight(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)/2
		if nums[m] == target {
			l = m + 1
		} else if nums[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	if r == -1 {
		return -1
	}
	if nums[r] == target {
		return r
	} else {
		return -1
	}
}
