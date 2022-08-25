package binaryTree

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	track := make([]int, 0)
	used := make([]bool, len(nums))
	var backTrack func()
	backTrack = func() {
		if len(track) == len(nums) {
			tmp := make([]int, len(track))
			for i := 0; i < len(track); i++ {
				tmp[i] = track[i]
			}
			res = append(res, tmp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}

			used[i] = true
			track = append(track, nums[i])
			backTrack()
			track = track[:len(track)-1]
			used[i] = false
		}
	}

	backTrack()
	return res
}
