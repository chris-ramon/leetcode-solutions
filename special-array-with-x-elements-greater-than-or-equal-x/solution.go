/*
Complexity Analysis
- Time: O(N * log N), where N is the size of the nums slice.
- Space: O(N), where N is the size of the nums slice.
*/

package main

import (
	"fmt"
	"sort"
)

func specialArray(nums []int) int {
	// size & > than N.
	sort.Slice(nums, func(a, b int) bool {
		return nums[a] < nums[b]
	})
	nums2 := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums2 = append(nums2, nums[i])
		}
	}
	size := len(nums2)
	max := nums2[len(nums2)-1]
	for i := 0; i < max; i++ {
		if i <= nums2[0] && i == size {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(specialArray([]int{3, 5}))
}
