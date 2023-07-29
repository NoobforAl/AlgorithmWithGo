package main

import (
	"fmt"
	"strconv"
)

/*
*
*
* this func for find Palindrome Number
* Wrong: this func <with> string convert
*
* args :
*   x int
*
* return value:
*   bool
*
*
 */
func isPalindromeConvert(x int) bool {
	// pars int value to string
	xStr := strconv.Itoa(x)

	/*
		/ check first and last value in string
		/ if not equal return true
		/ Order ( n )
	*/
	for i := 0; i < (len(xStr))/2; i++ {
		if xStr[i] !=
			xStr[len(xStr)-(1+i)] {

			return false
		}
	}

	return true
}

/*
*
*
* this func for find Palindrome Number
* Wrong: this func <without> string convert!
*
* args :
*   x int
*
* return value:
*   bool
*
*
 */
func isPalindrome(x int) bool {
	// check number is negative
	// can't negative number
	// is Palindrome
	if x < 0 {
		return false
	}

	// 0 <= x < 10
	// is Palindrome
	if x < 10 {
		return true
	}

	/*
		/ we make a copy of x to cx ( copy x )
		/ and revers x to res value
		/ number to number get last num
		/ from x and plus with res*10
	*/
	rev, cx := 0, x
	for x > 0 {
		l := x % 10
		rev = (rev * 10) + l
		x = x / 10
	}

	return cx == rev
}

func main() {
	// very simple test
	fmt.Println(isPalindromeConvert(1112111))

	// this without pars to string
	fmt.Println(isPalindrome(10))
}
