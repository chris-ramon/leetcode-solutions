// **Complexity Analysis**
// * Time: O(M + N), where M & N are the given slices.
// * Space: O(M + N), where M & N are the maps used to store the slice nums.

// ```
// Explanation:
// x: alice candies to give
// y: bob candies to give
// sum(a) - x + y = sum(b) - y + x
// - x - x = sum(b) - sum(a) - y - y
// - 2x = (sum(b) - sum(a)) - 2y
// 2y = (sum(b) - sum(a)) + 2x
// y = ((sum(b) - sum(a)) + 2x) / 2
// y = (sum(b) - sum(a) / 2) + x
// ```

func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	sumA := 0
	aliceNums := make(map[int]struct{})
	for _, n := range aliceSizes {
		sumA += n
		aliceNums[n] = struct{}{}
	}
	sumB := 0
	bobNums := make(map[int]struct{})
	for _, n := range bobSizes {
		sumB += n
		bobNums[n] = struct{}{}
	}
	for n, _ := range aliceNums {
		y := ((sumB - sumA) / 2) + n
		if _, found := bobNums[y]; found {
			return []int{n, y}
		}
	}
	return []int{}
}
