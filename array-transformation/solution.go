// **Complexity Analysis**
// * Time: O(max(N) - min(N)), where N is the length of the given arr.
// * Space: O(N), where N is the size of the queue during the breath-first search.

func transformArray(arr []int) []int {
	result := []int{}
	q := [][]int{arr}
	for len(q) > 0 {
		a := q[0]
		temp := make([]int, len(a))
		copy(temp, a)
		changed := false
		q = q[1:]
		size := len(a)
		for i := 0; i < size; i++ {
			if i == 0 || i == size-1 {
				continue
			}
			if a[i] > a[i-1] && a[i] > a[i+1] {
				temp[i]--
				changed = true
			}
			if a[i] < a[i-1] && a[i] < a[i+1] {
				temp[i]++
				changed = true
			}
		}
		if changed {
			q = append(q, temp)
		}
		result = a
	}
	return result
}
