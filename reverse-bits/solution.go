/*
Complexity Analysis
- Time: O(n) where n is the number of bits of the given num.
- Space: O(n) where n is the number used to store the result.
*/

// @lc code=start
func reverseBits(num uint32) uint32 {
	ret := uint32(0)
	power := uint32(31)
	for num != 0 {
		ret += (num & 1) << power
		num = num >> 1
		power -= 1
	}
	return ret
}
// @lc code=end

/*
TestCases
00000010100101000001111010011100\n11111111111111111111111111111101\n
*/
