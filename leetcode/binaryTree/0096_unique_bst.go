package binaryTree

func numTrees(n int) int {
	mem := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		mem[i] = make([]int, n+1)
	}
	var count func(lo, hi int) int
	count = func(lo, hi int) int {
		if lo > hi {
			return 1
		}

		if mem[lo][hi] != 0 {
			return mem[lo][hi]
		}
		res := 0
		for i := lo; i <= hi; i++ {
			res += count(lo, i-1) * count(i+1, hi)
		}
		mem[lo][hi] = res

		return res
	}

	return count(1, n)
}
