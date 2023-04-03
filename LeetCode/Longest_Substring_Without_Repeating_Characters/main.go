package main

import (
	"log"
)

/*
//* my first answer
//* this answer so bad
//* O(n^2)
func MaxSub(s string) int {
	var c int = 0
	check := make(map[rune]bool)
	for _, v := range s {
		if _, ok := check[v]; !ok {
			check[v] = false
			c++
		} else {
			return c
		}
	}
	return c
}

func lengthOfLongestSubstring(s string) int {
	var max int = 0
	for i := 0; i < len(s); i++ {
		if m := MaxSub(s[i:]); m > max {
			max = m
		}
	}
	return max
}
*/

/*
*
*
* find length Of Longest Substring
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
func lengthOfLongestSubstring(s string) int {
	var ans, l int = 0, 0
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if val, ok := m[s[i]]; ok {
			if val >= l {
				l = val + 1
			}
		}
		m[s[i]] = i
		if ans < i-l+1 {
			ans = i - l + 1
		}
	}
	return ans
}

func main() {
	// simple test
	if ans := lengthOfLongestSubstring("bbbb"); ans != 1 {
		log.Fatalf("Worn Answer! answer equal 3 not equal : %d", ans)
	}
	log.Println("Answer is True!")
}
