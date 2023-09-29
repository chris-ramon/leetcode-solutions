// Complexity Analysis
// Time: O(N), where N is the size of one row of the matrix.
// Space: O(1), constant number of variables used.

// @lc code=start
func diagonalSum(mat [][]int) int {
	primaryDiagonal := []int{0, 0}
	secondaryDiagonal := []int{0, len(mat[0])-1}
	result := 0
	for m := 0; m < len(mat); m++ {
		i, j := primaryDiagonal[0], primaryDiagonal[1]
		x, y := secondaryDiagonal[0], secondaryDiagonal[1]
		
		n1 := mat[i][j]
		n2 := mat[x][y]

		if i == x && j == y {
			result += n1
		} else {
			result += n1
			result += n2
		}

		primaryDiagonal[0]++
		primaryDiagonal[1]++

		secondaryDiagonal[0]++
		secondaryDiagonal[1]--
	}
	return result
}
// @lc code=end

/*
TestCases

[[1,2,3],[4,5,6],[7,8,9]]
[[1,1,1,1],[1,1,1,1],[1,1,1,1],[1,1,1,1]]
*/

