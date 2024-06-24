package main

/*
Complexity Analysis:
Time Complexity: O(g + s), where g + s are the length of the array g and s.
Space Complexity: O(g), where g is the size of the map.
*/
func findContentChildren(g []int, s []int) int {
	sMap := map[int]int{}
	result := 0

	for i := 0; i < len(s); i++ {
		if _, found := sMap[i]; found {
			sMap[i]++
		} else {
			sMap[i] = 1
		}
	}

	for i := 0; i < len(g); i++ {
		if v, found := sMap[g[i]]; found && v > 0 {
			sMap[i]--
			result++
		}
	}

	return result
}
