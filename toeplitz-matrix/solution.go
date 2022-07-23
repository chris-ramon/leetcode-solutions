// **Complexity Analysis**
// * Time: O(M * N), where M * N is the size of the given matrix.
// * Space: O(1), constant number of variables used.

func isToeplitzMatrix(matrix [][]int) bool {
	m := len(matrix)
	n := len(matrix[0])
	for i := 0; i < m-1; i++ {
		for j := 0; j < n-1; j++ {
			if matrix[i][j] != matrix[i+1][j+1] {
				return false
			}
		}
	}
	return true
}
