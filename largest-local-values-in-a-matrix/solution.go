package main

import "fmt"

func largestLocal(grid [][]int) [][]int {
	l := [][]int{}
	iOffsetLimit := len(grid) - 3 + 1 // 2
	jOffsetLimit := len(grid[0]) - 2
	iOffset := 0
  jOffset := 0
	for iOffset < iOffsetLimit {
		subL := []int{}
	  for jOffset < jOffsetLimit {
      fmt.Printf("jOffset: %+v\n", jOffset)
	  	for i := iOffset; i < 3; i++ {
	  		for j := jOffset; j < 3 + jOffset; j++ {
          fmt.Printf("j: %+v\n", j)
	  			subL = append(subL, grid[i][j])
	  		}
	  	}
	  	l = append(l, subL)
	  	jOffset++
	  }
    iOffset++
	}
  return l
}

func main() {
	fmt.Println(largestLocal([][]int{[]int{9, 9, 8, 1}, []int{5, 6, 2, 6}, []int{8, 2, 6, 4}, []int{6, 2, 2, 2}}))
}
