/*
Complexity Analysis
- Time: O(n), where n is the size of the given array nums.
- Space: O(n), where n is the size of the result.
*/


// @lc code=start
func getConcatenation(nums []int) []int {
	result := []int{}
	for _, num := range nums {
		result = append(result, num)
	}
	for _, num := range nums {
		result = append(result, num)
	}
	return result
}
// @lc code=end

/*
TestCases
[1,2,1]\n[1,3,2,1]\n[1,3,2,1,4,5,6]\n
*/