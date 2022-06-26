// **Complexity Analysis**
// * Time: O(N), where N is the given array's length.
// * Space: O(N), where N is the max size of the used hash table to store the array elements.

func uniqueOccurrences(arr []int) bool {
	m := make(map[int]int)
	for _, a := range arr {
		m[a] = m[a] + 1
	}
	n := make(map[int]struct{})
	for _, v := range m {
		if _, found := n[v]; found {
			return false
		}
		n[v] = struct{}{}
	}
	return true
}
