// **Complexity Analysis**
// - Time: O(N), where N is the size of the given strings.
// - Space: O(N), where N is the size of the map used to store given string characters.

func isIsomorphic(s string, t string) bool {
	charsMap := map[byte]byte{}
	tCharsMap := map[byte]struct{}{}
	for i := 0; i < len(s); i++ {
		if _, found := tCharsMap[t[i]]; !found {
			charsMap[s[i]] = t[i]
			tCharsMap[t[i]] = struct{}{}
		}
	}

	sReplaced := []byte{}
	for i := 0; i < len(s); i++ {
		sReplaced = append(sReplaced, charsMap[s[i]])
	}

	return string(sReplaced) == t
}
