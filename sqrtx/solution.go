// **Complexity Analysis**
// * Time: O(log N), where N is the given num.
// * Space: O(1), constant number of variables used.

func mySqrt(x int) int {
	if x < 2 {
		return x
	}
	left := 2
	right := x / 2
	for left <= right {
		mid := (left + right) / 2
		num := mid * mid
		if num > x {
			right = mid - 1
		} else if num < x {
			left = mid + 1
		} else {
			return mid
		}
	}
	return right
}
