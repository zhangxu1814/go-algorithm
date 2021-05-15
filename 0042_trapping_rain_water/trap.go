package trapping_rain_water

func MAX(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func MIN(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func trapBF(height []int) int {
	res := 0

	for i := 1; i < len(height)-1; i++ {
		//find the max value in the left
		ml := 0
		for j := i - 1; j >= 0; j-- {
			if height[j] > ml {
				ml = height[j]
			}
		}
		//find the max value in the right
		mr := 0
		for j := i + 1; j < len(height); j++ {
			if height[j] > mr {
				mr = height[j]
			}
		}

		if m := MIN(ml, mr); m > height[i] {
			res += m - height[i]
		}
	}

	return res
}

func trapDP(height []int) int {
	res, l := 0, len(height)
	//use 2 araay to record the max value of left and right for the current item
	mls, mrs := make([]int, l), make([]int, l)

	for i := 1; i < l-1; i++ {
		mls[i] = MAX(mls[i-1], height[i-1])
	}

	for i := l - 2; i > 0; i-- {
		mrs[i] = MAX(mrs[i+1], height[i+1])
	}

	for i := 1; i < l-1; i++ {
		if min := MIN(mls[i], mrs[i]); min > height[i] {
			res += min - height[i]
		}
	}

	return res
}

func trapDP2(height []int) int {
	if len(height) == 0 {
		return 0
	}

	res, maxLeft, maxRight, pointLeft, pointRight := 0, height[0], height[len(height)-1], 1, len(height)-2

	for pointLeft <= pointRight {
		if maxLeft < maxRight {
			if maxLeft > height[pointLeft] {
				res += maxLeft - height[pointLeft]
			} else {
				maxLeft = height[pointLeft]
			}
			pointLeft++
		} else {
			if maxRight > height[pointRight] {
				res += maxRight - height[pointRight]
			} else {
				maxRight = height[pointRight]
			}
			pointRight--
		}
	}

	return res
}
