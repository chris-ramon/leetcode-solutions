
func countOdds(low int, high int) int {
	result := 0
	for i := low; i <= high; i++ {
		if (i % 2) != 0 {
			result += 1
		}
	}
	return result
}
