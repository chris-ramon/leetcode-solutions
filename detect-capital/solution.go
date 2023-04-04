/*
Complexity Analysis
- Time: O(n), where n is the size of the given word.
- Space: O(1), constant number of variables used.
*/

// @lc code=start
func detectCapitalUse(word string) bool {
	re := regexp.MustCompile(`^[A-Z]{0,1}[a-z]+$`)
	re2 := regexp.MustCompile(`^[A-Z]+$`)
	return re.MatchString(word) || re2.MatchString(word)

}
// @lc code=end

/*
TestCases
"USA"\n"ABC"\n"abc"\n"Google"\n"FlaG"\n
*/
