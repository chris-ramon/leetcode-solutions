// @lc code=start
func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	for _, factor := range []int{2, 3, 5} {
		n = keepDividingWhenDivisible(n, factor)
	}
	return n == 1
}

func keepDividingWhenDivisible(n int, factor int) int {
	for n % factor == 0 {
		n /= factor
	}
	return n
}
// @lc code=end

/*
TestCases
6\n1\n14
8
*/