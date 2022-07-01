// Complexity Analysis
// - Time: O(N), where N is the size of given nums.
// - Space: O(N), where N is the map size used to store the nums count.
func majorityElement(nums []int) int {
	result := 0
	m := make(map[int]int)
	max := 0.0
	half := math.Round(float64(len(nums)) / 2.0)
	for i := 0; i < len(nums); i++ {
		if _, found := m[nums[i]]; found {
			m[nums[i]]++
		} else {
			m[nums[i]] = 1
		}
		if m[nums[i]] >= int(half) {
			max = math.Max(max, float64(m[nums[i]]))
			result = nums[i]
		}
	}
	return result
}
