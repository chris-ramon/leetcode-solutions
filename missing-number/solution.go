// **Complexity Analysis**
// * Time: O(N * logN), where N are the length of the given nums.
// * Space: O(1), constant number of variabled used.
func missingNumber(nums []int) int {
	left := 0
	right := len(nums) - 1
	mid := 0
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	for left <= right {
		mid = int(math.Round(float64(left+right) / 2.0))
		if nums[mid] > mid {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
