// **Complexity Analysis**
// * Time: O(N), where N is the size of the given num string.
// * Space: O(1), constant number of variables used.

func isStrobogrammatic(num string) bool {
	strobogrammaticMap := map[byte]byte{
		'6': '9',
		'9': '6',
		'0': '0',
		'1': '1',
		'8': '8',
	}
	p1 := 0
	p2 := len(num) - 1
	for p1 <= p2 {
		n1 := num[p1]
		n2 := num[p2]
		strobogrammatic := false
		if v, found := strobogrammaticMap[n1]; found {
			if v != n2 {
				return false
			}
			strobogrammatic = true
		}
		if v, found := strobogrammaticMap[n2]; found {
			if v != n1 {
				return false
			}
			strobogrammatic = true
		}
		if !strobogrammatic {
			return false
		}
		p1++
		p2--
	}
	return true
}
