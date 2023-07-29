package main

import "fmt"

/*
*
*
* validation string
*
* args :
*   str slices []string
*
* return value:
*   bool
*
*
 */
func isValid(str string) bool {
	// firs check if length
	// str is odd return false
	// else check if length str is zero
	// return true
	if len(str)%2 != 0 {
		return false
	} else if len(str) == 0 {
		return true
	}

	// make on map for check
	// symbol
	chars := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	// make one quest
	var q []rune

	for _, v := range str {

		// check we get close
		// symbol or not
		if k, ok := chars[v]; ok {

			// if len q is zero
			// this means we have
			// symbol close without
			// symbol open
			if len(q) < 1 {
				return false
			}

			// if symbol is different
			// like [ == }
			// return false
			if q[len(q)-1] != k {
				return false
			}

			// else pop last char in quest
			q = q[:len(q)-1]
			continue
		}

		// if not we append to quest
		q = append(q, v)
	}

	// now check len q
	// if is not zero
	// this means one open symbol
	// in quest!
	return len(q) == 0
}

func main() {
	// very simple test
	fmt.Println(isValid("([{[{()}]}])"))
}
