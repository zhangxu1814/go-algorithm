package array

import "sort"

func advantageCount(nums1 []int, nums2 []int) []int {
	type pair struct {
		i, v int
	}

	sort.Ints(nums1)

	tmp := make([]pair, len(nums2))
	res := make([]int, len(nums2))
	for i := 0; i < len(nums2); i++ {
		tmp[i] = pair{i, nums2[i]}
	}
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i].v < tmp[j].v
	})

	cur, l, r := len(tmp)-1, 0, len(nums1)-1
	for cur >= 0 {
		if nums1[r] > tmp[cur].v {
			res[tmp[cur].i] = nums1[r]
			r--
		} else {
			res[tmp[cur].i] = nums1[l]
			l++
		}
		cur--
	}

	return res
}
