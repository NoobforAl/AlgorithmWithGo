package main

import (
	"testing"
)

func TestCase1(t *testing.T) {
	nums := []int{1, 2, 3}
	ans := []int{1, 3}
	l := removeElement(nums, 2)
	if len(ans) != l {
		t.Error("Wrong Length!")
	}

	for i := 0; i < l; i++ {
		if ans[i] != nums[i] {
			t.Error("Wrong Answer!")
		}
	}
}

func TestCase2(t *testing.T) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	ans := []int{0, 1, 3, 0, 4}
	l := removeElement(nums, 2)

	if len(ans) != l {
		t.Error("Wrong Length!")
	}

	for i := 0; i < l; i++ {
		if ans[i] != nums[i] {
			t.Error("Wrong Answer!")
		}
	}
}
