func plusOne(digits []int) []int {
	num := ""
	for i := 0; i < len(digits); i++ {
		num += fmt.Sprintf("%d", digits[i])
	}

	n := new(big.Int)
	n.SetString(num, 10)

	r := new(big.Int)
	r.Add(n, big.NewInt(1))

	result := []int{}
	for _, di := range strings.Split(r.String(), "") {
		d, _ := strconv.Atoi(di)
		result = append(result, d)
	}

	return result
}


/*
TestCases
[1,2,3]\n[4,3,2,1]\n[9,9]\n[0]\n[7,2,8,5,0,9,1,2,9,5,3,6,6,7,3,2,8,4,3,7,9,5,7,7,4,7,4,9,4,7,0,1,1,1,7,4,0,0,6]\n
*/
