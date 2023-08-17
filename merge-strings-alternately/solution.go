// Complexity Analysis
// - Time: O(n), where n is the size of max(word1, word2).
// - Space: O(n), where n is the size of max(word1, word2), used to store the result.

func mergeAlternately(word1 string, word2 string) string {
	size := len(word1)
	if len(word2) > size {
		size = len(word2)
	}
	result := []byte{}
	for i := 0; i < size; i++ {
		if i < len(word1) {
			result = append(result, word1[i])
		}
		if i < len(word2) {
			result = append(result, word2[i])
		}
	}
	return string(result)
