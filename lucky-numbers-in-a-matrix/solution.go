/*
Complexity Analysis
- Time: O(n*m), where n and m are the columns and rows of the matrix.
- Time: O(n+m), where n and m are the columns and rows of the matrix.
*/

import "sort"

func luckyNumbers(matrix [][]int) []int {
	result := []int{}
	minMap := map[int]bool{}
	for i := 0; i < len(matrix); i++ {
		minValues := []int{}
		for j := 0; j < len(matrix[i]); j++ {
			v := matrix[i][j]
			minValues = append(minValues, v)
		}
		sort.Slice(minValues, func(i, j int) bool {
			return minValues[i] < minValues[j]
		})
		minMap[minValues[0]] = true
	}
	maxMap := map[int]bool{}
	for i := 0; i < len(matrix[0]); i++ {
		maxValues := []int{}
		for j := 0; j < len(matrix); j++ {
			v := matrix[j][i]
			maxValues = append(maxValues, v)
		}
		sort.Slice(maxValues, func(i, j int) bool {
			return maxValues[i] > maxValues[j]
		})
		maxMap[maxValues[0]] = true
	}
	for k, _ := range minMap {
		if _, ok := maxMap[k]; ok {
			result = append(result, k)
		}
	}
	return result
}

/*
TestCases
[[3,7,8],[9,11,13],[15,16,17]]\n[[1,10,4,2],[9,3,8,7],[15,16,17,12]]\n[[7,8],[1,2]]\n
*/
