// **Complexity Analysis**
// * Time: O(log N), where N is the given num.
// * Space: O(1), constant number of variables used.

func isPerfectSquare(num int) bool {
	if num < 2 {
		return true
	}
	low := 0
	high := num
	mid := num / 2
	for low < high {
		mid = (low + high) / 2
		switch {
		case mid*mid == num:
			return true
		case mid*mid > num:
			high = mid
		case low == mid:
			return false
		default:
			low = mid
		}
	}
	return false
}
