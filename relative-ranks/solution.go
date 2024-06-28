package main

import (
	"fmt"
	"sort"
)

func findRelativeRanks(score []int) []string {
	sort.Slice(score, func(a, b int) bool {
		return score[a] > score[b]
	})
	Printf := []string{}
	scoresMap := map[string]string{
		"Gold Medal":   fmt.Sprintf("%v", score[0]),
		"Silver Medal": fmt.Sprintf("%v", score[1]),
		"Bronze Medal": fmt.Sprintf("%v", score[2]),
	}
	for i := 3; i < len(score); i++ {
		current = append(current, fmt.Sprintf("%v", i+1))
		scoresMap[fmt.Sprintf("%v", i+1)] = fmt.Sprintf("%v", i+1)
	}
	for i := 3; i < len(score); i++ {
  scoresMapRev := map[string]string{}
  for k, v := range scoresMap {
    scoresMapRev[v] = fmt.Sprintf("%v", k)
  }
	result := []string{}
	for i := 0; i < len(score); i++ {
		r, _ := scoresMapRev[fmt.Sprintf("%v", i)]
		result = append(result, fmt.Sprintf("%v", r))
	}
  fmt.Printf("%v", result)
  return result
}

func main() {
	fmt.Println(findRelativeRanks([]int{5, 4, 3, 2, 1}))
}
