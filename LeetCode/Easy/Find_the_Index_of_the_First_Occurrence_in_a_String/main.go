package main

import (
	"log"
)

/*
*
*
* find index one string from
* another string
*
* args :
*   haystack, needle string
*
* return value:
*   int
*
*
 */
func strStr(haystack, needle string) int {
	// easy way is use strings pkg
	// import "strings"
	// return strings.Index(haystack, needle)

	// check string with len len(haystack)-len(needle)+1
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		// check if i to i+len(needle) equal needle
		if haystack[i:i+len(needle)] == needle {
			// return index
			return i
		}
	}

	// if exit loop return -1
	return -1
}

func main() {
	// very simple test
	if strStr("test0", "t0") != 3 {
		log.Fatal("Wrong Answer!")
	}

	log.Println("True Answer!")
}
