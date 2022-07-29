package array

type NumArray struct {
	preSum []int
}

//func Constructor(nums []int) NumArray {
//	tmp := make([]int, len(nums))
//	tmp[0] = nums[0]
//	for i := 1; i < len(nums); i++ {
//		tmp[i] = tmp[i-1] + nums[i]
//	}
//	return NumArray{preSum: tmp}
//}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.preSum[right]
	}
	return this.preSum[right] - this.preSum[left-1]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
