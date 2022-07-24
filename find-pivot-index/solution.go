// **Complexity Analysis**
// * Time: O(N), where N is the size of the given nums.
// * Space: O(1), constant number of variables used.

func pivotIndex(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	leftSum := 0
	for i := 0; i < len(nums); i++ {
		n := nums[i]
		if leftSum == (sum - n) {
			return i
		}
		leftSum += n
		sum -= n
	}
	return -1
}
