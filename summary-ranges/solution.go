// **Complexity Analysis**
// * Time: O(N), where N is the size of the given nums.
// * Space: O(N), where N is the size of the queue used to store the nums.

func summaryRanges(nums []int) []string {
	result := []string{}
	if len(nums) == 0 {
		return result
	}
	q := []int{nums[0]}
	addSummaryRange := func() {
		if len(q) == 1 {
			a := q[0]
			result = append(result, fmt.Sprintf("%d", a))
		}
		if len(q) > 1 {
			a := q[0]
			b := q[len(q)-1]
			result = append(result, fmt.Sprintf("%d->%d", a, b))
		}
	}
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] != 1 {
			addSummaryRange()
			q = []int{}
		}
		q = append(q, nums[i])
	}
	addSummaryRange()
	return result
}
