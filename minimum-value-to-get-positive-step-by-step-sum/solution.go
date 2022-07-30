// **Complexity Analysis**
// * Time: O(N), where N is the length of the given nums.
// * Space: O(1), constant number of variables used.

func minStartValue(nums []int) int {
	min := 0.0
	total := 0.0
	for i := 0; i < len(nums); i++ {
		total += float64(nums[i])
		min = math.Min(min, total)
	}
	return int(math.Abs(min) + 1)
}
