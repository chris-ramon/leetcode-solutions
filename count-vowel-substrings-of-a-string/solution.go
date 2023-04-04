/*
Complexity Analysis
- Time: O(n^2), where n is the size of the given word.
- Space: O(n), where n is the size of the given word.
*/

// @lc code=start
func countVowelSubstrings(word string) int {
	result := 0
	for i := 0; i < len(word); i++ {
		check := make(map[byte]int)
		for j := i; j < len(word); j++ {
			if strings.Contains("aeiou", string(word[j])) {
				check[word[j]] += 1
			} else {
				break
			}
			if len(check) == 5 {
				result++
			}
		}
	}
	return result
}
// @lc code=end

/*
TestCases
"aeiouu"\n"unicornarihan"\n"cuaieuouac"\n
*/