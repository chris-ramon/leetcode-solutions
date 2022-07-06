// **Complexity Analysis**
// * Time: O(N * log N), where N is the size of the given box types.
// * Space: O(N), where N is the size of the queue during the breadth-first search.

import "sort"

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
	result := 0
	sizeUsed := 0
	level := 0
	q := [][]int{boxTypes[0]}
	for len(q) > 0 {
		size := q[0][0]
		for i := 0; i < size; i++ {
			result += q[0][1]
			sizeUsed++
			if sizeUsed == truckSize {
				return result
			}
		}
		q = q[1:]
		level++
		if level < len(boxTypes) {
			q = append(q, boxTypes[level])
		}
	}
	return result
}
