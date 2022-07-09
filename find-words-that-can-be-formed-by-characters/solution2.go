func countCharacters(words []string, chars string) int {
	m := make(map[byte]int, len(chars))
	for i := 0; i < len(chars); i++ {
		m[chars[i]]++
	}
	result := 0
	for i := 0; i < len(words); i++ {
		size := len(words[i])
		result += size
		wordMap := make(map[byte]int, size)
		for j := 0; j < size; j++ {
			wordMap[words[i][j]]++
		}
		for letter, count := range wordMap {
			mCount, found := m[letter]
			if !found || count > mCount {
				result -= size
				break
			}
		}
	}
	return result
}