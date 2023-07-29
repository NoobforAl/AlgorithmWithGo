package main

import (
	"testing"
)

func TestCase1(t *testing.T) {
	nums := []int{1, 2, 3}
	ans := []int{1, 2, 3}
	l := removeDuplicates(nums)
	if len(ans) != l {
		t.Error("Wrong Length!")
	}

	for i := 0; i < l; i++ {
		if ans[i] != nums[i] {
			t.Error("Wrong Answer!")
			return
		}
	}
}

func TestCase2(t *testing.T) {
	nums := []int{1, 2, 3, 3, 3, 3, 4, 4, 4, 5, 5, 5, 5, 6, 6, 6, 6}
	ans := []int{1, 2, 3, 4, 5, 6}
	l := removeDuplicates(nums)

	if len(ans) != l {
		t.Error("Wrong Length!")
	}

	for i := 0; i < l; i++ {
		if ans[i] != nums[i] {
			t.Error("Wrong Answer!")
		}
	}
}
