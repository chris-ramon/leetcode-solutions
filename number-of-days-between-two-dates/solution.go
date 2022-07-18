// **Complexity Analysis**
// * Time: O(1), constant runtime.
// * Space: O(1), constant number of variables used.

import (
	"math"
	"time"
)
func daysBetweenDates(date1 string, date2 string) int {
	const layout = "2006-01-02"
	d1, _ := time.Parse(layout, date1)
	d2, _ := time.Parse(layout, date2)
	diff := d2.Sub(d1)
	return int(math.Abs(diff.Hours() / 24))
}
