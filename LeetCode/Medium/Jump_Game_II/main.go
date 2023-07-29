package main

import (
	"log"
)

/*
*
* jump
*
* args :
*   nums []int slice
*
* return value:
*   int
*
 */
func jump(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	currMax, nextMax := 0, 0
	steps, i := 0, 0

	for i <= currMax {
		for ; i <= currMax; i++ {
			nextMax = Max(nextMax, nums[i]+i)
			if nextMax >= n-1 {
				return steps + 1
			}
		}
		steps++
		currMax = nextMax
	}
	return steps
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// simple test
	if ans := jump([]int{0}); ans != 0 {
		log.Fatalf("Worn Answer! answer equal `0` not equal : %d", ans)
	}
	log.Println("Answer is True!")
}
