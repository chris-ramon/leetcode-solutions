// **Complexity Analysis**
// - Time: O(M * N), where M * N are the size of the given accounts.
// - Space: O(1), used constant number of variables.

func maximumWealth(accounts [][]int) int {
	result := 0.0
	for i := 0; i < len(accounts); i++ {
		sum := 0.0
		for j := 0; j < len(accounts[i]); j++ {
			sum += float64(accounts[i][j])
		}
		result = math.Max(result, sum)
	}
	return int(result)
}
