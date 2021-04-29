package sort

func BubbleSort(sequence []int) []int {
	l := len(sequence)
	//最坏情况下需要l-1轮，每轮排出一个最大值
	for i := 0; i < l; i++ {
		//判断当前待排序值是否已经有序
		sorted := true
		//第i轮已有i个值完成排序，剩余l-i个值待排序
		for j := 0; j < l-i-1; j++ {
			if sequence[j] > sequence[j+1] {
				sequence[j], sequence[j+1] = sequence[j+1], sequence[j]
				sorted = false
			}
		}

		if sorted {
			break
		}
	}

	return sequence
}
