package two_sum

func twoSum(nums []int, target int) []int {
	size := len(nums)
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func twoSumUsingMap(nums []int, target int) []int {
	size := len(nums)
	m := make(map[int]int)
	for i := 0; i < size; i++ {
		want := target - nums[i]
		if index, exist := m[want]; exist {
			return []int{index, i}
		}
		m[nums[i]] = i
	}
	return []int{}
}
