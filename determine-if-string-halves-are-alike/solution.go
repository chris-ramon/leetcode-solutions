// # Complexity
// - Time complexity: `O(N)`, where N is the size of the given string.
// - Space complexity: `O(1)`, constant number of variables used.

func halvesAreAlike(s string) bool {
	mid := len(s) / 2
	leftCount := 0
	rightCount := 0
	for idx, l := range s {
		switch {
		case idx < mid && isVowel(l):
			leftCount++
		case idx >= mid && isVowel(l):
			rightCount++
		}
	}
	return leftCount == rightCount
}

func isVowel(l rune) bool {
	switch l {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}
