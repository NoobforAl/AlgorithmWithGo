package main

import "testing"

func TestCase1(t *testing.T) {
	if !isPalindrome(112211) {
		t.Error("for number 112211 returned true!")
	}
}

func TestCase(t *testing.T) {
	if isPalindrome(1211) {
		t.Error("for number 1211 returned false!")
	}
}
