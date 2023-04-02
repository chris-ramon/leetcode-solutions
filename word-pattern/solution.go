func wordPattern(pattern string, s string) bool {
	patternMap := make(map[string]byte)
	visited := make([]bool, 26)
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}
	for i, word := range words {
		p, ok := patternMap[word]
		if !ok {
			if visited[pattern[i]-'a'] {
				return false
			}
			visited[pattern[i]-'a'] = true
			patternMap[word] = pattern[i]
			continue
		}
		if p != pattern[i] {
			return false
		}
	}
	return true
}