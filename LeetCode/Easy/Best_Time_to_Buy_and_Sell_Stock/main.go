package main

import (
	"log"
)

/*
*
*
* max Profit
*
* args :
*   prices []int slice
*
*
* return value:
*   int
*
*
 */
func maxProfit(p []int) int {
	buy, res := p[0], 0
	for i := 1; i < len(p); i++ {
		if buy > p[i] {
			buy = p[i]
		} else if p[i]-buy > res {
			res = p[i] - buy
		}
	}
	return res
}

func main() {
	// simple test
	if ans := maxProfit([]int{7, 6, 4, 3, 1}); ans != 0 {
		log.Fatalf("Worn Answer! answer equal `0` not equal : %d", ans)
	}
	log.Println("Answer is True!")
}
