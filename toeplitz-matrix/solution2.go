func isToeplitzMatrix(matrix [][]int) bool {
	m := len(matrix[0])
	n := len(matrix)
	nums := make(map[int]int)
	for i := 0; i < n; i++ {
		idx := n - i
		for j := 0; j < m; j++ {
			if n, found := nums[idx]; found {
				if n != matrix[i][j] {
					return false
				}
			} else {
				nums[idx] = matrix[i][j]
			}
			idx++
		}
	}
	return true
}