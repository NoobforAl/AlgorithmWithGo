package main

import "testing"

func TestCase1(t *testing.T) {
	if isValid("[") {
		t.Error("this test equal false!")
	}
}

func TestCase2(t *testing.T) {
	if !isValid("([])") {
		t.Error("this test equal true!")
	}
}
