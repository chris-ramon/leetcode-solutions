func reverseWords(s string) string {
	result := ""
	current := ""
	for _, c := range s {
		if c == ' ' {
			result += current + " "
			current = ""
		} else {
			current = string(c) + current
		}
	}
	result += current
	return result
}