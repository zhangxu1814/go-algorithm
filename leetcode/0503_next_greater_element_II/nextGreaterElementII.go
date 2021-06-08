package next_greater_element_II

func nextGreaterElementII(nums []int) []int {
	l := len(nums)
	result := []int{}
	for i, v := range nums {
		greater := -1
		for k := i + 1; k < l+i; k++ {
			if j := k % l; nums[j] > v {
				greater = nums[j]
				break
			}
		}
		result = append(result, greater)
	}

	return result
}

//TODO optimize this problem using stack
//maybe the % operation spent too much time
