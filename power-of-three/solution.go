// @lc code=start
func isPowerOfThree(n int) bool {
	if n < 1 {
		return false
	}
	for n % 3 == 0 {
		n /= 3
	}

	return n == 1
}
// @lc code=end

/*
TestCases
27\n0\n-1\n43046720\n
*/