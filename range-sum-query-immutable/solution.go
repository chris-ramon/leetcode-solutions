/*
TestCases
- Time: O(n), where n is the given nums.
- Space: O(n), where n is the size of the given nums.
*/

// @lc code=start
type NumArray struct {
	nums []int
}


func Constructor(nums []int) NumArray {
	return NumArray{
		nums: nums,
	}
}


func (this *NumArray) SumRange(left int, right int) int {
	result := 0
	sumEnabled := false
	for i := 0; i < len(this.nums); i++ {
		if i == left {
			sumEnabled = true
		}
		if sumEnabled {
			result += this.nums[i]
		}
		if i == right {
			sumEnabled = false
		}
	}
	return result
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
// @lc code=end



/*
TestCases
["NumArray","sumRange"]\n[[[-2,0,3,-5,2,-1]],[0,2]]\n
["NumArray","sumRange","sumRange","sumRange"]\n[[[-2,0,3,-5,2,-1]],[0,2],[2,5],[0,5]]\n
*/
