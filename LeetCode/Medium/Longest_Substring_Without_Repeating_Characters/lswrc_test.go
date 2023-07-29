package main

import (
	"testing"
)

func TestCase1(t *testing.T) {
	if ans := lengthOfLongestSubstring("abcabcbb"); ans != 3 {
		t.Fatalf("Worn Answer! answer equal 3 not equal : %d", ans)
	}
	t.Log("Answer is True!")
}

func TestCase2(t *testing.T) {
	if ans := lengthOfLongestSubstring("pwwkew"); ans != 3 {
		t.Fatalf("Worn Answer! answer equal 3 not equal : %d", ans)
	}
	t.Log("Answer is True!")
}
