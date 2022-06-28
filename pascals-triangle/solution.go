// **Complexity Analysis**
// * Time: O(N^2), where N is the number of given rows.
// * Space: O(N), where N is the size of the map used during the dp.

func generate(numRows int) [][]int {
	dp := make(map[int][]int)
	dp[0] = []int{1}
	dp[1] = []int{1, 1}

	for i := 2; i < numRows; i++ {
		nums := []int{}
		for j := 0; j < i-1; j++ {
			nums = append(nums, dp[i-1][j]+dp[i-1][j+1])
		}
		dp[i] = []int{1}
		dp[i] = append(dp[i], nums...)
		dp[i] = append(dp[i], 1)
	}

	result := [][]int{}
	for i := 0; i < numRows; i++ {
		result = append(result, dp[i])
	}
	return result
}
