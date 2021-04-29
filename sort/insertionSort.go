package sort

func InsertionSort(sequence []int) []int {
	l := len(sequence)
	//将第一个元素作为基值进行排序
	//第i轮前i个元素已经有序
	for i := 1; i < l; i++ {
		v := sequence[i]
		//每轮按顺序选择第一个待排序元素
		j := i - 1
		//将其按照从后往前的顺序进行比较
		for j >= 0 {
			//在有序序列中进行比较
			//若遇到不大于待排序元素的值，表明该值的后一个位置位置是待排序元素应当插入的位置，即可退出循环
			if sequence[j] <= v {
				break
			}
			//否则将该元素后移，继续向前寻找
			sequence[j+1] = sequence[j]
			j--
		}
		//将待排序元素插入
		sequence[j+1] = v
	}

	return sequence
}
