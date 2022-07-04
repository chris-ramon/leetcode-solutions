// Complexity Analysis
// - Time: O(N), where N is the size of the give slice.
// - Space: O(1), constant number of variables used.
func peakIndexInMountainArray(arr []int) int {
	for i := 1; i < len(arr)-1; i++ {
		left := arr[i-1]
		right := arr[i+1]
		if left < arr[i] && arr[i] > right {
			return i
		}
	}
	return -1
}
