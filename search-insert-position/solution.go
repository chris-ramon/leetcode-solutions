// **Complexity Analysis**
// * Time: O(log N), where N is the size of the given nums.
// * Space: O(1), constant number of variables used.

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		num := nums[mid]
		if target < num {
			right = mid - 1
		} else {
			left = mid + 1
		}
		if target == num {
			return mid
		}
	}
	return left
}
