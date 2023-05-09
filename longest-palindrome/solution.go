/*
Complexity Analysis
- Time: O(n), where n is the size of the given string.
- Space: O(1), constant number of variables used.
*/


func longestPalindrome(s string) int {
	count := [128]int{}
	for _, c := range s {
		count[c]++
	}
	result := 0
	for _, c := range count {
		result += c / 2 * 2
		if result % 2 == 0 && c % 2 == 1 {
			result++
		}
	}
	return result
}
