package main

func letterCombinations(digits string) []string {
	var result []string = []string{}

	if len(digits) == 0 {
		return result
	}

	dfs(digits, 0, &result)

	return result
}

var digitsMap map[rune][]rune = map[rune][]rune{
	'1': []rune{},
	'2': []rune{'a', 'b', 'c'},
	'3': []rune{'d', 'e', 'f'},
	'4': []rune{'g', 'h', 'i'},
	'5': []rune{'j', 'k', 'l'},
	'6': []rune{'m', 'n', 'o'},
	'7': []rune{'p', 'q', 'r', 's'},
	'8': []rune{'t', 'u', 'v'},
	'9': []rune{'w', 'x', 'y', 'z'},
}

func dfs(digits string, idx int, result *[]string) {
	if len(digits) <= idx {
		return
	}

	var letters []rune = digitsMap[rune(digits[idx])]

	if len(*result) == 0 {
		for _, l := range letters {
			letter := string(l)
			*result = append(*result, letter)
		}
		dfs(digits, idx+1, result)
		return
	}

	var newResult []string = []string{}

	for len(*result) > 0 {
		r := (*result)[0]
		*result = (*result)[1:]

		for _, j := range letters {
			letter := r + string(j)
			newResult = append(newResult, letter)
		}
	}

	*result = newResult

	dfs(digits, idx+1, result)
}
