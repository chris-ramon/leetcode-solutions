func isValid(s string) bool {
	if len(s) <= 1 {
		return false
	}
	stack := []rune{}
	for i := 0; i < len(s); i++ {
		c := rune(s[i])
		if c == '(' || c == '{' || c == '[' {
			stack = append(stack, rune(c))
		}
		if c == ')' || c == '}' || c == ']' {
			if len(stack) == 0 {
				return false
			}
			c2 := stack[len(stack)-1]
			if (c2 == '(' && c == ')') || (c2 == '{' && c == '}') || (c2 == '[' && c == ']') {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, rune(c))
			}
		}
	}
	return len(stack) == 0
}

/*
TestCases
{{}}
[]{}
()[]{}
{]
[()])
{}(())[(())]
{{{{{}}}}}()
*/
