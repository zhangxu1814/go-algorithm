package sort

func selectionSort(sequence []int) []int {
	l := len(sequence)
	//执行l-1轮，每轮从未排序元素中选出一个最小值交换至未排序序列最左
	for i := 0; i < l-1; i++ {
		min := i
		//遍历未排序序列，得到最小值的索引
		for j := i + 1; j < l; j++ {
			if sequence[j] < sequence[min] {
				min = j
			}
		}
		sequence[i], sequence[min] = sequence[min], sequence[i]
	}

	return sequence
}
