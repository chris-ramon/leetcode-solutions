// **Complexity Analysis**
// * Time: O(1), constant runtime.
// * Space: O(1), constant number of variables used.

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	return n&(n-1) == 0
}
