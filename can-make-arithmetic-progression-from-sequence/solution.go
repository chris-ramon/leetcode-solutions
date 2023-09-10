// Complexity Analysis
// Time: O(N * log N), where N is the sice of the given array.
// Space: O(1), constant number of variables used.

func canMakeArithmeticProgression(arr []int) bool {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
	diff := arr[0] - arr[1]
	for i := 1; i < len(arr); i++ {
		if arr[i-1] - arr[i] != diff {
			return false
		}
	}
	return true
}

/*
TestCases
[3,5,1]\n[1,2,4]
*/
