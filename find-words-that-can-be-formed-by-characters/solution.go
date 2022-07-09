// **Complexity Analysis**
// * Time: O(N), where N is the size of the given words slice.
// * Space: O(1), constant number of variables used.

func countCharacters(words []string, chars string) int {
	charsCount := [27]int{}
	for i := 0; i < len(chars); i++ {
		c := chars[i] % 96
		charsCount[c]++
	}
	result := 0
	for i := 0; i < len(words); i++ {
		size := len(words[i])
		result += size
		charsCountCopy := charsCount
		for j := 0; j < size; j++ {
			letter := words[i][j] % 96
			if charsCountCopy[letter] == 0 {
				result -= size
				break
			}
			charsCountCopy[letter]--
		}
	}
	return result
}
