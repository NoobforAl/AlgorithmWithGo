package main

import (
	"log"
	"reflect"
)

/*
*
*
* search Insert
*
* args :
*   digits []int slice
*
*
* return value:
*   []int slice
*
*
 */
func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return digits
	}

	var c int
	digits[len(digits)-1] += 1

	for i := len(digits) - 1; i > -1; i-- {
		sum := digits[i] + c
		if sum < 10 {
			digits[i] = sum
			return digits
		}
		c = sum / 10
		sum = sum % 10
		digits[i] = sum
	}

	if c != 0 {
		digits = append([]int{c}, digits...)
	}
	return digits
}

func main() {
	// simple test
	if ans := plusOne([]int{9}); !reflect.DeepEqual(ans, []int{1, 0}) {
		log.Fatalf("Worn Answer! answer equal `[1, 0]` not equal : %v", ans)
	}
	log.Println("Answer is True!")
}
