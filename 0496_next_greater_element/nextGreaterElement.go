package next_greater_element

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for i, v := range nums2 {
		m[v] = i
	}
	l2 := len(nums2)
	for i, v := range nums1 {
		j := m[v]
		if j == l2-1 {
			nums1[i] = -1
			continue
		}
		greater := -1
		for j++; j < l2; j++ {
			if nums2[j] > v {
				greater = nums2[j]
				break
			}
		}
		nums1[i] = greater
	}

	return nums1
}

//TODO optimize this problem using stack
