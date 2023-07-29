package main

import (
	"log"
)

func majorityElement(nums []int) int {
	last, cou := nums[0], 1

	for i := range nums {
		if last == nums[i] {
			cou++
		} else {
			cou--
			if cou == 0 {
				last, cou = nums[i], 1
			}
		}
	}

	return last
}

func main() {
	testArr := []int{2, 2, 3, 3, 3, 3, 2}

	if ans := majorityElement(testArr); ans != 3 {
		log.Fatalf("Answer Equal 3 not equal : %d", ans)
	}
	log.Println("All is Ok!")
}
