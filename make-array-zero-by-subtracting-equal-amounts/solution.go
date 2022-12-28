// # Complexity
// - Time complexity: `O(N * log N)`, where N is the size of the given nums.
// - Space complexity: `O(1)`, constant number of variables used.

import "sort"

func minimumOperations(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	result := 0
	for idx, num := range nums {
		if num != 0 {
			op(nums, idx)
			result++
		}
	}
	return result
}

func op(nums []int, idx int) {
	num := nums[idx]
	for i := idx; i < len(nums); i++ {
		nums[i] = nums[i] - num
	}
}
