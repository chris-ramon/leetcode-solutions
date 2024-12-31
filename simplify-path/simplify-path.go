/*
Complexity Analysis
- Time: O(n), n is the size of the given string.
- Space: O(n), n is the size of the result.
*/

package main

import (
	"fmt"
	"strings"
)

func simplifyPath(s string) string {
	pathValues := strings.Split(s, "/")
	values := []string{}
	for i := 0; i < len(pathValues); i++ {
		pathValue := pathValues[i]
		if pathValue != "" && pathValue != ".." {
			values = append(values, pathValue)
		}
		if (pathValue == ".." || pathValue == ".") && len(values) > 0 {
			values = values[:len(values)-1]
		}
	}
	result := strings.Join(values, "/")
	if len(result) > 0 && result[len(result)-1] == '/' {
		result = result[:len(result)-1]
	}
	result = fmt.Sprintf("/%v", result)
	return result
}

/*
TestCases
/home/"\n"/home//foo/"\n"/home/user/Documents/../Pictures"\n"/../"\n"/.../a/../b/c/../d/./"
"/a/b/../c"\n
"/home//foo"\n
"/../"\n
"/home/user/Documents/../Pictures"\n
*/
