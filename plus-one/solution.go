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

/*
TestCases
[1,2,3]\n[4,3,2,1]\n[9,9]\n[0]\n[7,2,8,5,0,9,1,2,9,5,3,6,6,7,3,2,8,4,3,7,9,5,7,7,4,7,4,9,4,7,0,1,1,1,7,4,0,0,6]\n
*/
