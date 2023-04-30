package main

import (
	"bytes"
	"log"
	"sort"

	"github.com/kr/pretty"
)

/*
*
*
* last Stone Weight
*
* args :
*   stones []int slice
*
*
* return value:
*   int
*
*
 */
func lastStoneWeight(stones []int) int {
	for len(stones) > 1 {
		sort.Ints(stones)
		x, y := stones[len(stones)-2], stones[len(stones)-1]
		stones = stones[:len(stones)-2]
		if x != y {
			stones = append(stones, y-x)
		}
	}
	if len(stones) == 1 {
		return stones[0]
	}
	return 0
}

func main() {
	// simple test
	var diff bytes.Buffer

	pretty.Fdiff(&diff, lastStoneWeight([]int{}), 0)

	if diff.String() != "" {
		log.Fatalf("Worn Answer!\n %s", diff.String())
	}
	log.Println("Answer is True!")
}
