package main

import (
	"log"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	s = strings.ToUpper(s)
	l, h := 0, len(s)-1

	for l < h {
		for l < h {
			if unicode.IsDigit(rune(s[l])) {
				break
			}
			if unicode.IsLetter(rune(s[l])) {
				break
			}
			l++
		}

		for l < h {
			if unicode.IsDigit(rune(s[h])) {
				break
			}
			if unicode.IsLetter(rune(s[h])) {
				break
			}
			h--
		}

		if s[l] != s[h] {
			return false
		}
		h--
		l++
	}
	return true
}

func main() {
	testStr := " "
	if !isPalindrome(testStr) {
		log.Fatalln("This is Palindrome")
	}
	log.Println("It's Ok!")
}
