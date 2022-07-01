// Complexity Analysis
// - Time: O(N), where N is the size of given nums.
// - Space: O(1), constant number of variabled used. 
func majorityElement(nums []int) int {
	result := 0
	count := 0
	for _, num := range nums {
		if count == 0 {
			result = num
		}
		if num == result {
			count += 1
		} else {
			count -= 1
		}
	}
	return result
}
