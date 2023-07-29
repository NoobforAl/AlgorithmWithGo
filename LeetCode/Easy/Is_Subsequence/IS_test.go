package main

import "testing"

func TestCase1(t *testing.T) {
	if !isSubsequence("abc", "ahbgdc") {
		t.Error("this test equal true!")
	}
}

func TestCase(t *testing.T) {
	if isSubsequence("axc", "ahbgdc") {
		t.Error("this test equal false!")
	}
}
