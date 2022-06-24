// **Complexity Analysis**
// * Time: O(N), where N is the size of the given words.
// * Space: O(N) where N is the size of the queue used store the words.

func validWordAbbreviation(word string, abbr string) bool {
	queue := []rune{}
	for _, c := range word {
		queue = append(queue, c)
	}
	prevNum := ""
	for i := 0; i < len(abbr); i++ {
		c := abbr[i]
		if _, err := strconv.Atoi(string(c)); err != nil {
			if prevNum == "" {
				if len(queue) == 0 {
					return false
				}
				wordChar := queue[0]
				queue = queue[1:]
				if wordChar == rune(c) {
					continue
				} else {
					return false
				}
			} else {
				if strings.HasPrefix(prevNum, "0") {
					return false
				}
				num, err := strconv.Atoi(prevNum)
				if err != nil {
					return false
				}
				if num+1 > len(queue) {
					return false
				}
				for n := 0; n < num; n++ {
					queue = queue[1:]
				}
				prevNum = ""
				i--
			}
		} else {
			if prevNum == "" {
				prevNum = string(c)
			} else {
				prevNum += string(c)
			}
		}
	}
	if prevNum != "" && strings.HasPrefix(prevNum, "0") {
		return false
	}
	if prevNum != "" {
		n, err := strconv.Atoi(prevNum)
		if err != nil {
			return false
		}
		return len(queue) == n
	}
	return len(queue) == 0
}


/*
TestCases

"internationalization"
"i12iz4n"
"apple"
"a2e"
"substitution"
"s0ubstitut"
"substitution"
"s55n"
"internationalization"
"i5a11o1"
"a"
"01"
"hi"
"1hi"
"hi"
"01i"
"a"
"ab"
*/
