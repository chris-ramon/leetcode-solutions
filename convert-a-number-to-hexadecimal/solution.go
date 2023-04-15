/*
Complexity Analysis
- Time: O(1), constant runtime.
- Space: O(1), constant number of variables used.
*/

func toHex(num int) string {
	return fmt.Sprintf("%x", uint32(num))
}
