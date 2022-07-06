// **Complexity Analysis**
// * Time: O(N), where N is the size of the given slice.
// * Space: O(1), constant number of variables used.

func arraySign(nums []int) int {
	negativeNumbers := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			return 0
		}
		if nums[i] < 0 {
			negativeNumbers++
		}
	}
	if negativeNumbers % 2 == 0 {
		return 1
	}
	return -1
}

/*
TestCases
[1]\n
[-1]\n
[-1,-2,-3,-4,3,2,1]\n
[1,5,0,2,-3]\n
[-1,1,-1,1,-1]\n
[9,72,34,29,-49,-22,-77,-17,-66,-75,-44,-30,-24]\n
*/
