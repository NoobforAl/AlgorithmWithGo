package main

import (
	"testing"
)

func TestCase1(t *testing.T) {
	if ans := longestPalindrome("cbbd"); ans != "bb" {
		t.Fatalf("Worn Answer! answer equal 'bb' not equal : %s", ans)
	}
	t.Log("Answer is True!")
}

func TestCase2(t *testing.T) {
	if ans := longestPalindrome("babad"); ans != "aba" && ans != "bab" {
		t.Fatalf("Worn Answer! answer equal 4 not equal : %s", ans)
	}
	t.Log("Answer is True!")
}
