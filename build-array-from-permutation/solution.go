// **Complexity Analysis**
// * Time: O(N), where N is the size of the given nums.
// * Space: O(1), constant number of variables used.

func buildArray(nums []int) []int {
	mask := 1<<10 - 1
	for i := 0; i < len(nums); i++ {
		num := nums[nums[i]] & mask
		nums[i] = num<<10 | nums[i]
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = nums[i] >> 10
	}
	return nums
}
