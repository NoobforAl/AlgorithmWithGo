package main

import (
	"log"
)

/*
*
*
* climb Stairs
*
* args :
*   n int
*
* return value:
*   int
*
*
 */
func climbStairs(n int) int {
	var hash = map[int]int{
		0: 0,
		1: 1,
		2: 2,
		3: 3,
	}

	var f func(n int) int
	f = func(n int) int {
		if v, ok := hash[n]; ok {
			return v
		}
		val := f(n-2) + f(n-1)
		hash[n] = val
		return val
	}
	return f(n)
}

func main() {
	// simple test
	if ans := climbStairs(2); ans != 2 {
		log.Fatalf("Worn Answer! answer equal `2`` not equal : %v", ans)
	}
	log.Println("Answer is True!")
}
