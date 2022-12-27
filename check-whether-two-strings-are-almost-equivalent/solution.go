// # Complexity
// - Time complexity: `O(M + N)`, where M & N are the sizes of the given words.
// - Space complexity: `O(1)`, constant number of variables used.

func checkAlmostEquivalent(word1 string, word2 string) bool {
	w1 := make([]int, 26)
	for _, c := range word1 {
		w1[c%97]++
	}
	w2 := [26]int{}
	for _, c := range word2 {
		w2[c%97]++
	}
	for i := 0; i < 26; i++ {
		if math.Abs(float64(w1[i]-w2[i])) > 3 {
			return false
		}
	}
	return true
}
