package main

import (
	"log"
)

/*
*
*
* search Insert
*
* args :
*   nums []int slice
*	target int
*
* return value:
*   int
*
* O(log n)
*
 */
func searchInsert(nums []int, target int) int {
	//* Binary Search
	var low, high = 0, len(nums)
	for low != high {
		mid := (low + high) / 2
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}

func main() {
	// simple test
	if ans := searchInsert([]int{1, 3, 5, 6}, 5); ans != 2 {
		log.Fatalf("Worn Answer! answer equal 2 not equal : %d", ans)
	}
	log.Println("Answer is True!")
}
