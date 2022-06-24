// **Complexity Analysis**
// * Time: O(Log N), where N is the give nums length.
// * Space: O(1), inline solution.

func search(nums []int, target int) int {
	p1 := 0
	p2 := len(nums) - 1
	for p1 <= p2 {
		if nums[p1] == target {
			return p1
		}
		if nums[p2] == target {
			return p2
		}
		p3 := (p1 + p2) / 2
		mid := nums[p3]
		if target == mid {
			return p3
		}
		if target > mid {
			p1 = p3 + 1
		} else {
			p2 = p3 - 1
		}
	}
	return -1
}

/*
TestCases
[-1,0,3,5,9,12]
9
[-1,0,3,5,9,12]
2
[10]
10
[10,11]
1
[5]
-5
*/
