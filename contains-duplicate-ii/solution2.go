func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int][]int)
	for i := 0; i < len(nums); i++ {
		if _, found := m[nums[i]]; found {
			m[nums[i]] = append(m[nums[i]], i)
		} else {
			m[nums[i]] = []int{i}
		}
	}
	fK := float64(k)
	for _, idxs := range m {
		if len(idxs) >= 2 {
			for i := 0; i < len(idxs); i++ {
				for j := 1; j < len(idxs); j++ {
					if i == j {
						continue
					}
					if math.Abs(float64(idxs[i]-idxs[j])) <= fK {
						return true
					}
				}
			}
		}
	}
	return false
}