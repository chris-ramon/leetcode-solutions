// **Complexity Analysis**
// * Time: O(N), where N is the size of the given string.
// * Space: O(1), constant number of variabled used.

func removePalindromeSub(s string) int {
	isPalindrome := func(s string) bool {
		left := 0
		right := len(s) - 1
		for left < right {
			if s[left] == s[right] {
				left++
				right--
			} else {
				return false
			}
		}
		return true
	}
	if len(s) == 0 {
		return 0
	}
	if isPalindrome(s) {
		return 1
	}
	return 2
}
