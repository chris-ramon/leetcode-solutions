/*
Complexity Analysis:
- Time: O(n), where n is the given num.
- Space: O(1), constant number of variables used. 
*/

func isPowerOfFour(n int) bool {
	if n < 1 {
		return false
	}
	for n % 4 == 0 {
		n /= 4
	}
	return n == 1
}

/*
TestCases
16\n5\n1\n
*/
