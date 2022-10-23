func bitwiseComplement(n int) int {
	num := fmt.Sprintf("%b", n)
	result := ""
	for i := 0; i < len(num); i++ {
		n := num[i]
		if n == '1' {
			result += "0"
		} else {
			result += "1"
		}
	}
	// fmt.Printf("result: %+v\n", result)
	r, _ := strconv.ParseInt(result, 2, 64)
	return int(r)
}
