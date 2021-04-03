func decodeString(s string) string {
    result := ""
    stack := []rune{}
    for _, st := range s {
        stack = append(stack, st)
        decode := true
        subStr := ""
        subNum := ""
        if st == ']' {
            stack = stack[:len(stack)-1]
            for decode && len(stack) > 0 {
                str := stack[len(stack)-1]
                stack = stack[:len(stack)-1]
                fmt.Printf("str: %c \n", str)
                if str == '[' {
                    continue
                }
                if str == ']' {
                    break
                }
                if unicode.IsNumber(str) {
                    subNum = string(str) + subNum
                    continue
                }
                subStr = string(str) + subStr
            }
            nums, err := strconv.Atoi(subNum)
            fmt.Printf("nums: %v \n", nums)
            if err != nil {
                fmt.Println(err)
                break
            }
            for i := 0; i < nums; i++ {
                result = result + string(subStr)
            }
        }
    }
    for _, st := range stack {
        result += string(st)
    }
    return result
}
