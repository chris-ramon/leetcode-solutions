// **Complexity Analysis**
// * Time: O(target * log(target)).
// * Space: O(target * log(target)).

func racecar(target int) int {
	q := [][]interface{}{{"A", "", 0, 1}}
	visitedMap := make(map[string]struct{})
	visitedKey := func(position, speed int) string {
		return fmt.Sprintf("%d-%d", position, speed)
	}
	visited := func(position, speed int) bool {
		_, found := visitedMap[visitedKey(position, speed)]
		return found
	}
	setVisited := func(position, speed int) {
		visitedMap[visitedKey(position, speed)] = struct{}{}
	}
	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			inst, seq, p, s := q[0][0].(string), q[0][1].(string), q[0][2].(int), q[0][3].(int)
			q = q[1:]
			setVisited(p, s)
			if p == target {
				return len(seq)
			}
			if inst == "A" {
				p += s
				s *= 2
			}
			if inst == "R" {
				if s > 0 {
					s = -1
				} else {
					s = 1
				}
			}
			if !visited(p, s) {
				q = append(q, []interface{}{"A", seq + "A", p, s})
				if ((p+s) > target && s > 0) || ((p+s) < target && s < 0) {
					q = append(q, []interface{}{"R", seq + "R", p, s})
				}
			}
		}
	}
	return 0
}
