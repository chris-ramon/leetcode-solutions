// **Complexity Analysis**
// * Time: O(N * log N), where N is the size of the given nums.
// * Space: O(N), where N is the size of the map used to store the nums values.

import "sort"

func frequencySort(nums []int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	sort.Slice(nums, func(i, j int) bool {
		iCount := m[nums[i]]
		jCount := m[nums[j]]
		if iCount == jCount {
			return nums[i] > nums[j]
		}
		return iCount < jCount
	})
	return nums
}
