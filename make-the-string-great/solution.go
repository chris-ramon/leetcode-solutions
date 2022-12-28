// # Complexity
// - Time complexity: `O(N * M)`, where N is the size of the given string and M is the number of pairs in the string.
// - Space complexity: `O(M)`, where M is the number of pairs in the string.

func makeGood(s string) string {
	return dfs(s)
}

func dfs(s string) string {
	if s == "" || len(s) == 1 {
		return s
	}
	nextStr := (*string)(nil)
	for i := 0; i < len(s)-1; i++ {
		curr := string(s[i])
		next := string(s[i+1])
		if curr != next && (strings.ToUpper(curr) == next || strings.ToLower(curr) == next) {
			n := s[:i] + s[i+2:]
			nextStr = &n
			break
		}
	}
	if nextStr == nil {
		return s
	}
	return dfs(*nextStr)
}
