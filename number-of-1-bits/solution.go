// **Complexity Analysis**
// * Time: O(1), constant runtime up-to 32 bits.
// * Space: O(1), constant number of variables used.

func hammingWeight(num uint32) int {
	result := 0
	mask := uint32(1)
	for i := 0; i < 32; i++ {
		if num&mask != 0 {
			result++
		}
		mask = mask << 1
	}
	return result
}

/*
TestCases
00000000000000000000000000001011
00000000000000000000000010000000
11111111111111111111111111111101
*/
