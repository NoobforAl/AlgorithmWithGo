package main

import "testing"

func TestCase1(t *testing.T) {
	list1 := convertToList([]int{})
	list2 := convertToList([]int{})

	if checkSorted(mergeTwoLists(list1, list2),
		&[]int{}) {

		t.Error("this equal empty!")
	}
}

func TestCase2(t *testing.T) {
	list1 := convertToList([]int{0})
	list2 := convertToList([]int{})

	if checkSorted(mergeTwoLists(list1, list2),
		&[]int{0}) {

		t.Error("this equal [0]!")
	}
}
