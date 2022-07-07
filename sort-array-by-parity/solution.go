// **Complexity Analysis**
// * Time: O(N * log N), where N is the size of the given nums.
// * Space: O(1), constant number of variables used.

func sortArrayByParity(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i]%2 == 0
	})
	return nums
}
