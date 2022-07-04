// **Complexity Analysis**
// * Time: O(log N), where N is the size of the given slice.
// * Space: O(1), constant number of variables used.

func peakIndexInMountainArray(arr []int) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] > arr[mid+1] {
			right = mid - 1
		} else {
			left = mid + 1
		}
		fmt.Printf("mid: %+v\n", mid)
	}
	return left
}
