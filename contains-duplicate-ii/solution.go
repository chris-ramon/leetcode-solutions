// **Complexity Analysis**
// * Time: O(N), where N is the size of the given nums.
// * Space: O(N), where N is the size of the map used to store the nums values.

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, found := m[nums[i]]; found {
			return true
		} else {
			m[nums[i]] = i
			if len(m) > k {
				delete(m, nums[i-k])
			}
		}
	}
	return false
}
