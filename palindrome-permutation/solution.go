/*
Complexity Analysis
- Time: O(n), where n is the size of the given string.
- Space: O(n), where n is the size of the map used to store the result.
*/

// @lc code=start
func canPermutePalindrome(s string) bool {
	h := make(map[rune]bool, len(s))
	for _, c := range s {
		if _, found := h[c]; found {
			delete(h, c)
		} else {
			h[c] = true
		}
	}
	if len(h) < 2 {
		return true
	}
	return false
}
// @lc code=end

/*
TestCases
"code"\n"aab"\n"carerac"\n"aba"\n
*/