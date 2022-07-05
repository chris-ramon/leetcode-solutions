// **Complexity Analysis**
// * Time: O(N), where N is the size of the given string.
// * Space: O(1), constant number of variables used.

// ```
// Explanation:
// "A" 	-> (26 * 0) + 1 = 1
// "B" 	-> (26 * 0) + 2 = 2
// "AA" 	-> (26 * 1) + 1 = 27
// "BB" 	-> (26 * 2) + 2 = 56
// "ZY" 	-> (26 * 26) + 26 = 701
// "AAA" 	-> (26 * 26 * 1) + (26 * 1) + 1 = 703
// "AAA" 	-> (26^2 * 1) + (26^1 * 1) + (26^0 * 1) = 703
// "BB" 	-> (26^1 * 2) + (26^0 * 2) = 56
// "AB" 	-> (26^1 * 1) + (26^0 * 2) = 28
// ```

func titleToNumber(columnTitle string) int {
	result := 0.0
	s := strings.ToLower(columnTitle)
	size := len(s)
	for i := 0; i < size; i++ {
		num := s[i] % 96
		n := size - 1 - i
		result += math.Pow(float64(26), float64(n)) * float64(num)
	}
	return int(result)
}

/*
TestCases
"A"\n
"BB"\n
"AB"\n
"AA"\n
"AAA"\n
"ZY"\n
"ZZZ"\n
*/
