package main

import "testing"

func TestCase1(t *testing.T) {
	ans := twoSum([]int{1, 2, 3, 4}, 10)
	if ans != nil {
		t.Error("this test answer is equal nil!")
	}
}

func TestCase2(t *testing.T) {
	ans := twoSum([]int{1, 2, 3, 4}, 3)
	if ans == nil {
		t.Error("this test answer is not equal nil!")
	}
}
