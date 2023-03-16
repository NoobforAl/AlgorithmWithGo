package main

import "testing"

func TestCase1(t *testing.T) {

	if !arrayStringsAreEqual(
		[]string{"aaa", "bbb"},
		[]string{"aa", "abb", "b"},
	) {
		t.Error("this test answer is true!")
	}
}

func TestCase2(t *testing.T) {
	if arrayStringsAreEqual(
		[]string{"aca", "bb"},
		[]string{"a", "ab", "b"},
	) {
		t.Error("this test answer is false!")
	}
}
