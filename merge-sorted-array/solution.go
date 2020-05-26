package main

/*
[1,2,3,0,0,0]
3
[2,5,6]
3

[4,5,6,0,0,0]
3
[1,2,3]
3

[4,0,0,0,0,0]
1
[1,2,3,5,6]
5

[1,2,4,5,6,0]
5
[3]
1

[-1,0,0,3,3,3,0,0,0]
6
[1,2,2]
3
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	var p1 int = m - 1
	var p2 int = n - 1
	var p3 int = m + n - 1

	for p1 >= 0 || p2 >= 0 {
		switch {
		case p1 < 0:
			nums1[p3] = nums2[p2]
			p3 -= 1
			p2 -= 1
		case p2 < 0:
			nums1[p3] = nums1[p1]
			p3 -= 1
			p1 -= 1
		default:
			if nums1[p1] > nums2[p2] {
				nums1[p3] = nums1[p1]
				p1 -= 1
			} else {
				nums1[p3] = nums2[p2]
				p2 -= 1
			}
			p3 -= 1
		}
	}
}
