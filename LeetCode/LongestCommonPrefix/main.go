package main

import "fmt"

/*
*
*
* find longest common prefix
*
* args :
*   nums slices []string
*
* return value:
*   string
*
*
 */
func longestCommonPrefix(strs []string) string {
	// if length strs in 0 zero
	// return empty string
	if len(strs) == 0 {
		return ""
	}

	// this for find word
	// common prefix
	var flag bool = true

	// response string
	var res string = ""

	// loop to first value
	// in slice and check
	// char <i> in prefix
	// other words or not
	for i := range strs[0] {

		// get one char form
		// firs word in strs
		// and get char <i>
		key := strs[0][i]

		// loop to all strs
		for _, v := range strs {

			// if length one word in strs
			// equal zero return empty string
			if len(v) == 0 {
				return ""
			}

			// if length one less i
			// flag false
			// and break loop
			// and in down
			// return res
			if i > len(v)-1 {
				flag = false
				break
			}

			// if find one word equal
			// key flag true and continue
			if key == v[i] {
				flag = true
				continue
			}

			// if not flag false
			// and break words loop
			flag = false
			break
		}

		// if flag is true
		// append key
		if flag {
			res += string(key)
			flag = false

			// else return res
		} else {
			return res
		}
	}

	// if all words like
	// first word
	// exit loop and
	// return res
	return res
}

func main() {
	// very simple test
	fmt.Println(longestCommonPrefix(
		[]string{"flower", "flow", "flight"},
	))
}
