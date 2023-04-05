/*
Complexity Analysis
- Time: O(n), where n is the size of the given string.
- Space: O(n), where n is the size of the given string.
*/

// @lc code=start
func lengthOfLastWord(s string) int {
	words := strings.Split(strings.TrimSpace(s), " ")
	last := words[len(words)-1]
	return len(last)
}
// @lc code=end

/*
TestCases
"Hello World"\n"   fly me   to   the moon  "\n"luffy is still joyboy"\n
*/