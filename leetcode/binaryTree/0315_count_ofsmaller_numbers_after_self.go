package binaryTree

func countSmaller(nums []int) []int {
	type pair struct {
		index, val int
	}
	tmp, count, pairs := make([]pair, len(nums)), make([]int, len(nums)), make([]pair, len(nums))
	for i := 0; i < len(nums); i++ {
		pairs[i] = pair{
			index: i,
			val:   nums[i],
		}
	}
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
			tmp[i] = pairs[i]
		}
		l, r = lo, mid+1
		for i := lo; i <= hi; i++ {
			if l > mid {
				pairs[i] = tmp[r]
				r++
			} else if r > hi {
				pairs[i] = tmp[l]
				count[pairs[i].index] += r - mid - 1
				l++
			} else if tmp[l].val > tmp[r].val {
				pairs[i] = tmp[r]
				r++
			} else {
				pairs[i] = tmp[l]
				count[pairs[i].index] += r - mid - 1
				l++
			}
		}
	}

	sort(0, len(nums)-1)
	return count
}
