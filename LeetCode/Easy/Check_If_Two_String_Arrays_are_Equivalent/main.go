package main

import (
	"fmt"
	"strings"
)

/*
*
*
* check two slice words is equal?
*
* args:
*	word1: []string ( slice )
*	word2: []string ( slice )
*
* return value:
*	bool
*
*
 */
func arrayStringsAreEqual(
	word1 []string,
	word2 []string,
) bool {
	// comperes word 1 and 2
	// in slice to one string
	// and check tow string
	// equal or not! ( true or false )
	return strings.Join(word1, "") ==
		strings.Join(word2, "")
}

func main() {
	// very simple test
	fmt.Println(arrayStringsAreEqual(
		[]string{"aa", "bb"},
		[]string{"a", "ab", "b"},
	))
}
