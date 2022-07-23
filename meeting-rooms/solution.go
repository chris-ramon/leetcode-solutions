// **Complexity Analysis**
// * Time: O(N * log N), where N is the size of the intervals.
// * Space: O(1), constant number of variables used.

import (
	"sort"
)

func canAttendMeetings(intervals [][]int) bool {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i][1] > intervals[i+1][0] {
			return false
		}
	}
	return true
}

/*
TestCases
[[0,30],[5,10],[15,20]]\n[[7,10],[2,4]]\n[[1,10000],[1,2]]\n[[1,2]]\n
*/
