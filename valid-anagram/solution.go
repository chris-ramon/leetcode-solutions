// **Complexity Analysis**
// * Time: O(N), where N is the size of the given string s.
// * Space: O(N), where N is the size of the map used to store the letter counts of the given string s.

func isAnagram(s string, t string) bool {
	sMap := make(map[byte]int, len(s))
	for i := 0; i < len(s); i++ {
		sMap[s[i]]++
	}
	for i := 0; i < len(t); i++ {
		count, found := sMap[t[i]]
		if count == 0 {
			return false
		}
		if found {
			sMap[t[i]]--
		} else {
			return false
		}
	}
	for _, v := range sMap {
		if v > 0 {
			return false
		}
	}
	return true
}
