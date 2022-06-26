// **Complexity Analysis**
// * Time: O(N), where N is the length of the given string.
// * Space: O(N), where N is the size of the slice to store the groups.

func countBinarySubstrings(s string) int {
	groups := make([]int, len(s))
	t := 0
	groups[0] = 1
	for i := 1; i < len(s); i++ {
		if s[i-1] == s[i] {
			groups[t] += 1
		} else {
			t += 1
			groups[t] = 1
		}
	}
	result := 0.0
	for i := 1; i < len(groups); i++ {
		result += math.Min(float64(groups[i-1]), float64(groups[i]))
	}
	return int(result)
}
