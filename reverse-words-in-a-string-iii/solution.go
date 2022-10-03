// **Complexity Analysis**

// * Time Complexity: O(N), where N is the size of the given string.
// * Space Complexity: O(N), where N is the size of the strings builder used to store the result.

func reverseWords(s string) string {
	result := strings.Builder{}
	word := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			result.Write(word)
			result.Write([]byte(" "))
			word = []byte{}
		} else {
			word = append([]byte{s[i]}, word...)
		}
	}
	result.Write(word)
	return result.String()
}
