// **Complexity Analysis**
// * Time: O(N/2), where N is the given num.
// * Space: O(N), where N is the size of the slice to store the result.

func sumZero(n int) []int {
	result := []int{}
	mid := n / 2
	for i := 1; i < mid+1; i++ {
		result = append(result, i)
		result = append(result, i*-1)
	}
	if n%2 != 0 {
		result = append(result, 0)
	}
	return result
}
