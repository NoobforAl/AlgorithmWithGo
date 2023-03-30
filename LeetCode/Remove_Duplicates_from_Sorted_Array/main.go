package main

import (
	"log"
	"math"
)

/*
*
*
* This func for remove Duplicate Numbers
*
* args:
*	nums []int slice
*
* return:
*	int
*
*
 */
func removeDuplicates(nums []int) int {
	// first make index for slice
	// without duplicates nums
	// make last value
	// for first try set max int
	var index, last int = 0, math.MaxInt
	for _, n := range nums {
		// check if n != last value
		// set n value to nums[index]
		// index ++ and
		// last value = n
		if n != last {
			nums[index] = n
			index++
			last = n
		}
	}
	return index
}

func main() {
	// very simple test
	nums := []int{1, 1, 1}
	ans := []int{1}
	for i := 0; i < removeDuplicates(nums); i++ {
		if ans[i] != nums[i] {
			log.Fatal("Wrong Answer!")
		}
	}

	log.Println("True Answer!")
}
