// **Complexity Analysis**
// * Time: O(1), finite set of roman numerals.
// * Space: O(1), constant numer of variables.

func romanToInt(s string) int {
	symbols := map[string]int{
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}
	if result, found := symbols[s]; found {
		return result
	}
	result := 0
	for i := len(s) - 1; i >= 1; i-- {
		curr := string(s[i])
		prev := string(s[i-1])
		if val, found := symbols[prev+curr]; found {
			result += val
			if i > 1 {
				result -= symbols[prev]
			}
		} else {
			if i == 1 {
				result += symbols[prev]
				result += symbols[curr]
			} else {
				result += symbols[curr]
			}
		}
	}
	return result
}

/*
TestCases
"III"
"IV"
"X"
"M"
"MV"
"III"
"XXVII"
"LVIII"
"MCMXCIV"
"CMLII"
"MCDLXXVI"
*/
