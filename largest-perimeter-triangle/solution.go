// **Complexity Analysis**
// * Time: O(N * Log N), where N is the size of the given nums.
// * Space: O(1), constant number of variables used.
// * Note: Triangle inequality theorem.

func largestPerimeter(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	for i := 0; i < len(nums)-2; i++ {
		a := nums[i]
		b := nums[i+1]
		c := nums[i+2]
		if a < (b + c) {
			return a + b + c
		}
	}
	return 0
}

/*
TestCases
[2,1,2]\n[1,2,1]\n[1,1,1]\n[1,2,3,99,100,101]
*/
