package main

import (
	"fmt"
	"math"
)

//	      [0,1,0,2,1,0,1,3,2,1,2,1]
// p0        ^
// p1          ^
// p2            ^
func trap(height []int) int {
	p0 := 1 
	p1 := 1 
	p2 := 1 
	trappedWater := 0
	size := len(height)
	for p1 < size && p2 < size {
    fmt.Printf("p0: %v, p1: %v, p2: %v\n", p0, p1, p2)
		// 1. Put p2 pointer in place.
		if height[p0] == height[p2] || height[p0] > height[p2] {
			p2++
      continue
		}

		// 2. Move pointer p1.
		for i := p0; i < p2-p0; i++ {
      p1 += p0 + 1
			min := math.Min(float64(height[p0]), float64(height[p2]))
			trappedWater += int(min) - height[p1]
      fmt.Printf("i: %v, tw: %+v, min:%+v\n", i, trappedWater, min)
		}

		// 3. Reset Pointers.
    p0 = p2
    p1 = p2
    p2 = p2 + 1
	}
	return trappedWater
}

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2}))
}
