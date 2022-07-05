// Complexity Analysis
// - Time: O(M * N), where M & N are the rows & columns of the matrix.
// - Space: O(M * N), where M & N are the rows & columns of the matrix.
func transpose(matrix [][]int) [][]int {
	result := [][]int{}
	for i := 0; i < len(matrix[0]); i++ {
		result = append(result, make([]int, len(matrix)))
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			num := matrix[i][j]
			result[j][i] = num
		}
	}
	return result
}
