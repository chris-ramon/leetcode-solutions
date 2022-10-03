func reverseWords(s string) string {
	result := []string{}
	words := strings.Split(s, " ")
	for i := 0; i < len(words); i++ {
		word := words[i]
		reversedWord := []byte{}
		for j := len(word) - 1; j >= 0; j-- {
			reversedWord = append(reversedWord, word[j])
		}
		result = append(result, string(reversedWord))
	}
	return strings.Join(result, " ")
}