package main

import (
	"bytes"
	"log"
	"math"

	"github.com/kr/pretty"
)

/*
*
*
* reverse
*
* args :
*   x int
*
* return value:
*   int
*
*
 */
func reverse(x int) (r int) {
	f := false
	if x < 0 {
		f = true
	}

	x = int(math.Abs(float64(x)))
	for x != 0 {
		r = r*10 + x%10
		x /= 10
		if r >= math.MaxInt32 ||
			r <= math.MinInt32 {
			return 0
		}
	}

	if f {
		return 0 - r
	}
	return
}

func main() {
	var err bytes.Buffer
	pretty.Fdiff(&err, reverse(111111111111), 0)
	if err.String() != "" {
		log.Fatal(err.String())
	}
	log.Println("All OK!")
}
