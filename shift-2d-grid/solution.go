// **Complexity Analysis**
// * Time: O(m * n), where m & n are the number of rows & columns in the given matrix.
// * Space: O(m * n * k), where m & n are the new matrix created to store the new values on each iteration (k times).

func shiftGrid(grid [][]int, k int) [][]int {
	shift := func(grid [][]int) [][]int {
		result := [][]int{}
		m := len(grid)
		n := len(grid[0])
		for i := 0; i < m; i++ {
			result = append(result, make([]int, n))
		}
		for i := 0; i < m; i++ {
			for j := 0; j < n-1; j++ {
				result[i][j+1] = grid[i][j]
			}
			if i < m-1 {
				result[i+1][0] = grid[i][n-1]
			}
		}
		result[0][0] = grid[m-1][n-1]
		return result
	}
	result := grid
	for i := 0; i < k; i++ {
		result = shift(result)
	}
	return result
}
