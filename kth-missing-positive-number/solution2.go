// Complexity Analysis
// - Time: O(N), where N is the given array length.
// - Space: O(1), constant number of variabled used.
func findKthPositive(arr []int, k int) int {
	kthMissing := 0
	current := 1
	for i := 0; i < len(arr)+k; i++ {
		if i < len(arr) {
			if arr[i] != current {
				kthMissing++
				i--
			}
		} else {
			kthMissing++
		}
		if kthMissing == k {
			return current
		}
		current++
	}
	return 0
}

/*
TestCases
[1,2,3]\n1
[2,3,4,7,11]\n5
[1,2,3,4,5]\n10
[1,2,3,4]\n2
[1]\n5
*/
