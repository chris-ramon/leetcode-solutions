package main

import "fmt"

func timeRequiredToBuy(tickets []int, k int) int {
	// k, person at index.
	times := tickets[k]
	result := 0
	for i := 0; i < times; i++ {
		for i := 0; i < len(tickets); i++ {
			if tickets[i] != 0 {
				result++
				tickets[i]--
			}
		}
	}
	return result
}

func main() {
	// fmt.Println(timeRequiredToBuy([]int{2, 3, 2}, 2))
	// fmt.Println(timeRequiredToBuy([]int{2, 3, 2}, 2))
	fmt.Println(timeRequiredToBuy([]int{84, 49, 5, 24, 70, 77, 87, 8}, 3))
}
