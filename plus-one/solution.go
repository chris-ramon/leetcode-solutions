// **Complexity Analysis**
// * Time: O(N), where N is the length of the given digits slice.
// * Space: O(N), where N is the length of digits slice used to store the result.

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		n := digits[i]
		if n < 9 {
			digits[i]++
			return digits
		} else {
			digits[i] = 0
		}
	}
	digits = append([]int{1}, digits...)
	return digits
}
