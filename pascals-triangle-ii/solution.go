/*
Complexity Analysis
- Time: O(n), where n is the size of the given rows.
- Space: O(n), where n size of the map used to store the dp result.
*/
// @lc code=start
func getRow(rowIndex int) []int {
	dp := make(map[int][]int)
	dp[0] = []int{1}
	dp[1] = []int{1, 1}
	dp[2] = []int{1, 2, 1}
	for i := 3; i <= rowIndex; i++ {
		prev := dp[i-1]
		r := []int{1}
		for j := 1; j < len(prev); j++ {
			p1 := prev[j-1]
			p2 := prev[j]
			r = append(r, p1 + p2)
		}
		r = append(r, 1)
		dp[i] = r
	}
	return dp[rowIndex]
}
// @lc code=end

/*
TestCases
3\n0\n1\n4\n5\n6\n
*/