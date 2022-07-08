// **Complexity Analysis**
// * Time: O(N), where N is the size of the nums2.
// * Space: O(N), where N is the size of the map, used to store the nums from nums2.

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	nextGreaterMap := make(map[int]int, len(nums2))
	stack := []int{}
	for i := 0; i < len(nums2); i++ {
		num := nums2[i]
		for len(stack) > 0 && stack[len(stack)-1] < num {
			n := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			nextGreaterMap[n] = num
		}
		stack = append(stack, num)
	}
	result := []int{}
	for i := 0; i < len(nums1); i++ {
		num := nums1[i]
		nextGreaterNum, found := nextGreaterMap[num]
		if found {
			result = append(result, nextGreaterNum)
		} else {
			result = append(result, -1)
		}
	}
	return result
}

/*
TestCases
[4,1,2]\n[1,3,4,2]
[2,4]\n[1,2,3,4]
[10]\n[1,2,3,4,10]
*/
