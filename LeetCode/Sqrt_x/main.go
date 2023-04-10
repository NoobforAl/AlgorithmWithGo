package main

import (
	"log"
)

/*
*
*
* my Sqrt
*
* args :
*   x int
*
*
* return value:
*   int
*
*
 */
func mySqrt(x int) int {
	var i int
	for ; i*i <= x; i++ {
		if i*i == x {
			return i
		}
	}
	return i - 1
}

func main() {
	// simple test
	if ans := mySqrt(1); ans != 1 {
		log.Fatalf("Worn Answer! answer equal `1` not equal : %d", ans)
	}
	log.Println("Answer is True!")
}

// Better Answer
func BetterAnswer(x int) int {
	low, high := 0, x
	for low <= high {
		mid := (low + high) / 2

		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return high
}
