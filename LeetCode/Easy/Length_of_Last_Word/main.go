package main

import (
	"log"
)

/*
*
*
* find Length of Last Word
*
* args :
*   s string
*
* return value:
*   int
*
* O(n)
*
 */
func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}

	c, f := 0, false
	for i := len(s) - 1; i > -1; i-- {
		if s[i] != ' ' {
			c++
			f = true
		} else if f && s[i] == ' ' {
			return c
		}
	}
	return c
}

func main() {
	// simple test
	if ans := lengthOfLastWord("Hello World"); ans != 5 {
		log.Fatalf("Worn Answer! answer equal 3 not equal : %d", ans)
	}
	log.Println("Answer is True!")
}
