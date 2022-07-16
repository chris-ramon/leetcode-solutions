func hammingWeight(num uint32) int {
	n := fmt.Sprintf("%b", num)
	result := 0
	for i := 0; i < len(n); i++ {
		if n[i] == '1' {
			result++
		}
	}
	return result
}