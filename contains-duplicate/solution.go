// **Complexity Analysis**
// * Time: O(N), where N is the lenght of the given nums.
// * Space: O(N), where N is the max size of the map used to store the nums.

func containsDuplicate(nums []int) bool {
	m := make(map[int]struct{})
	for _, n := range nums {
		if _, found := m[n]; found {
			return true
		}
		m[n] = struct{}{}
	}
	return false
}
