import "sort"
func highFive(items [][]int) [][]int {
	students := make(map[int][]int)
	for i := 0; i < len(items); i++ {
		id, score := items[i][0], items[i][1]
		students[id] = append(students[id], score)
	}
	studentIDs := []int{}
	for id, _ := range students {
		studentIDs = append(studentIDs, id)
	}
	sort.Slice(studentIDs, func(i, j int) bool {
		return studentIDs[i] < studentIDs[j]
	})
	result := [][]int{}
	for _, id := range studentIDs {
		scores := students[id]
		sort.Slice(scores, func(i, j int) bool {
			return scores[i] > scores[j]
		})
		sum := 0
		for i := 0; i < 5; i++ {
			sum += scores[i]
		}
		topFiveAverage := sum / 5
		result = append(result, []int{id, topFiveAverage})
	}
	return result
}