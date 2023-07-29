package main

import "testing"

func TestCase1(t *testing.T) {
	if romanToInt("") != 0 {
		t.Error("This answer equal zero!")
	}
}

func TestCase2(t *testing.T) {
	if romanToInt("LVIII") != 58 {
		t.Error("This answer equal 58!")
	}
}
