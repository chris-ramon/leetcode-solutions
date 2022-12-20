// # Complexity Analysis
// - Time complexity: O(n), where n is the size of the given num.
// - Space complexity: O(n), where n is the size of the slice to store the result.

import (
	"fmt"
	"strconv"
	"strings"
)

func maximum69Number(num int) int {
	n := fmt.Sprintf("%d", num)
	result := strings.Split(n, "")
	for idx, c := range result {
		if c == "6" {
			result[idx] = "9"
			break
		}
	}
	r, _ := strconv.Atoi(strings.Join(result, ""))
	return r
}

/*
TestCases
9669\n9996\n9999\n9\n6\n66
*/
