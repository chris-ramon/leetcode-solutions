/*
TestCases
- Time: O(n), where n is the max(str1, str2).
- Space: O(n), where n is the size of the callstack during the dfs.
*/

// @lc code=start
func dfs(x int, y int) int {
	if y == 0 {
		return x
	} else {
		return dfs(y, x % y)
	}
}

func gcdOfStrings(str1 string, str2 string) string {
	size1 := str1 + str2
	size2 := str2 + str1
	if size1 != size2 {
		return ""
	}
	size := dfs(len(str1), len(str2))
	return str1[0: size]
}
// @lc code=end