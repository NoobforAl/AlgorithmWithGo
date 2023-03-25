package main

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	nums := []int{1, 2, 3}
	if removeDuplicates(&nums) != 3 &&
		reflect.DeepEqual(nums, []int{1, 2, 3}) {

		t.Error("Wrong Answer!")
	}
}

func TestCase2(t *testing.T) {
	nums := []int{1, 2, 3, 3, 3, 3, 4, 4, 4, 5, 5, 5, 5, 6, 6, 6, 6}
	if removeDuplicates(&nums) != 6 &&
		reflect.DeepEqual(nums, []int{1, 2, 3, 4, 5, 6}) {

		t.Error("Wrong Answer!")
	}
}
