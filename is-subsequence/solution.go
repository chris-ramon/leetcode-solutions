// **Complexity Analysis**
// 
// * Time: O(N), where N is the size of the given string s.
// * Space: O(1), constant number of variables used.

func isSubsequence(s string, t string) bool {
	p1, p2 := 0, 0
	for p1 < len(s) && p2 < len(t) {
		l1, l2 := s[p1], t[p2]
		if l1 == l2 {
			p1++
			p2++
		} else {
			p2++
		}
	}
	return p1 == len(s)
}
