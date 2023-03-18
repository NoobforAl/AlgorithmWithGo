package main

import "fmt"

/*
*
*
* this func for is Subsequence
* for two string
*
* args :
*   s string
*   t string
*
* return value:
*   bool
*
*
 */
func isSubsequence(s, t string) bool {
	// if length s is zero return true
	// empty s alway in t
	if len(s) == 0 {
		return true
	}

	// if length t is zero return false
	// because nothing can find in empty
	if len(t) == 0 {
		return false
	}

	// loop to main string < t >
	for i, j := 0, 0; i < len(t); i++ {

		// if t[i] is equal s[j]
		// go to next string in s[j]
		if t[i] == s[j] {
			j++
		}

		// if j equal to len s
		// return true
		// this show find all
		// s string in t!
		if j > len(s)-1 {
			return true
		}
	}

	// if exit from loop return false
	// this show not match or found s in t string
	return false
}

func main() {
	// very simple test
	fmt.Println(isSubsequence("abc", "ahbgdc"))
}
