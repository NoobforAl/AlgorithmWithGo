package main

import (
	"log"
)

/*
*
*
* This func for remove custom
* Value from slice
*
* args:
*	nums []int slice
*	val int
*
* return:
*	int
*
*
 */
func removeElement(nums []int, val int) int {
	// counter for length
	// slice without val number
	var i int = 0
	for _, n := range nums {

		// check if n != vla
		// last index (i)
		// get value n
		// and i = i + 1
		if n != val {
			nums[i] = n
			i++
		}
	}
	return i
}

func main() {
	// very simple test
	nums := []int{1, 1, 1}
	ans := []int{}
	for i := 0; i < removeElement(nums, 1); i++ {
		if ans[i] != nums[i] {
			log.Fatal("Wrong Answer!")
		}
	}

	log.Println("True Answer!")
}
