/*
# Complexity Analysis:
- Time: O(n * log n), where n is the size of the given salary slice.
- Space: O(1), constant number of variables used.
*/

func average(salary []int) float64 {
	sort.Slice(salary, func(i, j int) bool {
		return salary[i] > salary[j]
	})
	sum := 0.0
	count := 0.0
	for i := 1; i < len(salary)-1; i++ {
		sum += float64(salary[i])
		count++
	}
	return sum / count
}
