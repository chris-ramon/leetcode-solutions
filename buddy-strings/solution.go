// **Complexity Analysis**
// * Time: O(N), where N is the size of the given string s & goal.
// * Space: O(1), constant number of variables used.

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	firstDiff := -1
	secondDiff := -1
	lettersCount := [26]int{}
	for i := 0; i < len(s); i++ {
		if s[i] != goal[i] {
			if firstDiff == -1 {
				firstDiff = i
			} else if secondDiff == -1 {
				secondDiff = i
			} else {
				return false
			}
		}
		letter := s[i] % 96
		lettersCount[letter]++
	}
	if s == goal {
		for _, count := range lettersCount {
			if count > 1 {
				return true
			}
		}
		return false
	}
	if firstDiff == -1 || secondDiff == -1 {
		return false
	}
	return s[firstDiff] == goal[secondDiff] && s[secondDiff] == goal[firstDiff]
}
