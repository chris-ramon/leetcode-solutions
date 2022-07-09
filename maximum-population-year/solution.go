// **Complexity Analysis**
// * Time: O(N * log N), where N is the size of the given logs.
// * Space: O(N), where N is the size of the map, used to store the year's population.

import "sort"
func maximumPopulation(logs [][]int) int {
	years := []int{}
	m := make(map[int]int, len(logs))
	for i := 0; i < len(logs); i++ {
		b, d := logs[i][0], logs[i][1]
		m[b]++
		m[d]--
		if b == d {
			years = append(years, b)
		} else {
			years = append(years, b)
			years = append(years, d)
		}
	}
	sort.Slice(years, func(i, j int) bool {
		return years[i] < years[j]
	})
	start := years[0]
	end := years[len(years)-1]
	population := 0
	resultYear := start
	resultPopulation := 0
	for y := start; y <= end; y++ {
		if _, found := m[y]; found {
			population += m[y]
		}
		if population > resultPopulation {
			resultYear = y
			resultPopulation = population
		}
	}
	return resultYear
}
