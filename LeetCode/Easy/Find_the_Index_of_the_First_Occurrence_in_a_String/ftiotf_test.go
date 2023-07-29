package main

import "testing"

func TestCase1(t *testing.T) {
	if strStr("sadbutsad", "sad") != 0 {
		t.Error("Wrong Answer!")
	}
}

func TestCase2(t *testing.T) {
	if strStr("leetcode", "t0") != -1 {
		t.Error("Wrong Answer!")
	}
}
