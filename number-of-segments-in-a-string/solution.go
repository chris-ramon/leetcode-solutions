/*
Complexity Analysis
- Time: O(n), where n is the size of the given string.
- Space: O(1), constant number of variables used.
*/

func countSegments(s string) int {
	s = strings.Trim(s, " ")
	if s == "" {
		return 0
	}

	r := regexp.MustCompile("[\\s]+").FindAllStringIndex(s, -1)
	return len(r) + 1
}

/*
TestCases
"Hello, my name is John"\n"Hello"\n"a,      "\n
"love live! mu'sic forever"\n
*/
