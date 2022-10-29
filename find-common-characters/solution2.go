func commonChars(words []string) []string {
	wordsMaps := []map[byte]int{}
	for i := 0; i < len(words); i++ {
		word := words[i]
		charactersMap := make(map[byte]int)
		for j := 0; j < len(word); j++ {
			c := word[j]
			charactersMap[c]++
		}
		wordsMaps = append(wordsMaps, charactersMap)
	}
	// fmt.Printf("wordsMaps: %+v\n", wordsMaps)
	result := []string{}
	if len(wordsMaps) > 0 {
		w1 := wordsMaps[0]
		for k1, v1 := range w1 {
			v := float64(v1)
			foundAll := false
			for i := 1; i < len(wordsMaps); i++ {
				wordsMap := wordsMaps[i]
				if currentV, found := wordsMap[k1]; found {
					v = math.Min(v, float64(currentV))
				} else {
					break
				}
				if i == len(wordsMaps)-1 {
					foundAll = true
				}
			}
			if foundAll {
				for i := 0; i < int(v); i++ {
					result = append(result, string(k1))
				}
			}
		}
	}
	return result
}
