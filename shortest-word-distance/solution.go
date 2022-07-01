// **Complexity Analysis**
// * Time: O(N), where N is the lenght of the given words dict.
// * Space: O(1), constan number of variables are used.

func shortestDistance(wordsDict []string, word1 string, word2 string) int {
	w1 := -1.0
	w2 := -1.0
	result := float64(len(wordsDict))
	for i := 0; i < len(wordsDict); i++ {
		if wordsDict[i] == word1 {
			w1 = float64(i)
		}
		if wordsDict[i] == word2 {
			w2 = float64(i)
		}
		if w1 != -1.0 && w2 != -1.0 {
			result = math.Min(result, math.Abs(w1-w2))
		}
	}
	return int(result)
}

/*
TestCases
["a","b","b","c","a","z","z","z","c"]\n"a"\n"c"
["practice", "makes", "perfect", "coding", "makes"]\n"coding"\n"practice"
["practice", "makes", "perfect", "coding", "makes"]\n"makes"\n"coding"
*/
