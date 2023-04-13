/*
Complexity Anaylsis
- Time: O(max(n)), where n is the max size between s || t.
- Space: O(max(n)), where n is max size of the maps.
*/

// @lc code=start
func findTheDifference(s string, t string) byte {
	countS := map[rune]int{}
	countT := map[rune]int{}
	for _, v := range s {
		countS[v]++
	}
	for _, v := range t {
		countT[v]++
	}
	for k, countT := range countT {
		if countS, found := countS[k]; !found || countT != countS {
			return byte(k)
		}
	}
	var b byte
	return b
}
// @lc code=end

/*
TestCases
"abcd"\n"abcde"\n""\n"y"\n"a"\n"aa"\n
*/
