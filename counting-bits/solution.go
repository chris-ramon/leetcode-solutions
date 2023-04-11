/*
Complexity Analysis
- Time: O(n), where n is the given num n.
- Space: O(n), where n is the array used to store the result.
*/


// @lc code=start
func countBits(n int) []int {
	result := []int{}
	for i := 0; i <= n; i++ {
		bits := fmt.Sprintf("%b", i)
		count := 0
		for j := 0; j < len(bits); j++ {
			if bits[j] == '1' {
				count++
			}
		}
		result = append(result, count)
	}
	return result
}
// @lc code=end

/*
TestCases
2\n5\n1\n
*/
