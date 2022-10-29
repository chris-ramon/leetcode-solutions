// **Complexiy Analysis**
// * Time: O(N), where N is the size of the given nums.
// * Space: O(N), where N is the size of the slice used to store the result.


func shuffle(nums []int, n int) []int {
	result := []int{}
	p1 := 0
	p2 := n
	for p2 < len(nums) {
		n1, n2 := nums[p1], nums[p2]
		result = append(result, n1, n2)
		p1++
		p2++
	}
	return result
}

/*
TestCases
[2,5,1,3,4,7]\n3\n[1,2,3,4,4,3,2,1]\n4\n[1,1,2,2]\n2
*/
