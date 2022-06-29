// **Complexity Analysis**
// * Time: O(N), where N is the size of the given string.
// * Space: O(N-D), where D is the number of duplicates.

func removeDuplicates(s string) string {
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && s[i] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}
