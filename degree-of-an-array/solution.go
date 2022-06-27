// **Complexity Analysis**
// * Time: O(N), where N is the size of the given nums.
// * Space: O(N), where N is the size of the map used to store the nums.

type Num struct {
	Count     int
	LowerIdx  int
	HigherIdx int
}

func findShortestSubArray(nums []int) int {
	m := make(map[int]*Num)
	degree := 0
	if len(nums) > 0 {
		degree = 1
	}
	for i, num := range nums {
		if n, found := m[num]; found {
			n.Count++
			n.HigherIdx = i
			if n.Count > degree {
				degree = n.Count
			}
		} else {
			m[num] = &Num{Count: 1, LowerIdx: i, HigherIdx: i}
		}
	}
	min := math.MaxFloat64
	for _, num := range m {
		if degree == num.Count {
			min = math.Min(min, float64(num.HigherIdx-num.LowerIdx)+1)
		}
	}
	return int(min)
}
