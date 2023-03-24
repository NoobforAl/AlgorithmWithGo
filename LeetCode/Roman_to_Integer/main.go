package main

import "fmt"

/*
*
*
* roman To Int
*
* args :
*   s slices []string
*
* return value:
*   int
*
*
 */
func romanToInt(s string) int {
	// make response value
	var res int = 0

	// map symbol with value
	symbol := map[byte]int{
		'I': 1, 'V': 5,
		'X': 10, 'L': 50,
		'C': 100, 'D': 500,
		'M': 1000,
	}

	// map combine symbol with value
	combineSymbol := map[string]int{
		"IV": 4, "IX": 9,
		"XL": 40, "XC": 90,
		"CD": 400, "CM": 900,
	}

	// for every char in string
	// check char
	for i := 0; i < len(s); i++ {

		/*
			/ if i+1 exist in length s
			/ return res + value of s[i]
			/ else
			/ check if s[i] ( C, I, X )
			/ go to check combine symbol map
			/ if found res + value of map
			/ else
			/ res + value of symbol
		*/
		if i+1 >= len(s) {
			return res + symbol[s[i]]
		} else if s[i] == 'C' {
			if v, ok := combineSymbol["C"+string(s[i+1])]; ok {
				i, res = i+1, res+v
				continue
			}
		} else if s[i] == 'I' {
			if v, ok := combineSymbol["I"+string(s[i+1])]; ok {
				i, res = i+1, res+v
				continue
			}
		} else if s[i] == 'X' {
			if v, ok := combineSymbol["X"+string(s[i+1])]; ok {
				i, res = i+1, res+v
				continue
			}
		}

		res = res + symbol[s[i]]
	}

	return res
}

func main() {
	// simple test
	fmt.Println(romanToInt("DCXXI") == 621)
}
