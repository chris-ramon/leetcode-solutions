// **Complexity Analysis**
// * Time: O(1), constant number of iterations.
// * Space: O(1), constant number of variables used.

func bitwiseComplement(n int) int {
	if n == 0 {
		return 1
	}
	current := n
	bit := 1
	for current != 0 {
		n = n ^ bit
		bit = bit << 1
		current = current >> 1
	}
	return n
}
