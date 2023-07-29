package main

import "testing"

func TestCase1(t *testing.T) {
	ans := longestCommonPrefix([]string{"flower", "flow", "flight"})
	if ans != "fl" {
		t.Error("this test answer is equal fl!")
	}
}

func TestCase2(t *testing.T) {
	ans := longestCommonPrefix([]string{"dog", "racecar", "car"})
	if ans != "" {
		t.Error("this test answer is equal empty!")
	}
}
