package main

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	tree := sortedArrayToBST([]int{-10, -3, 0, 5, 9})
	ans := reflect.DeepEqual(BFS(tree), []int{0, -3, 9, -10, 5})
	if !ans {
		t.Error("Wrong Answer!")
	}
}

func TestCase2(t *testing.T) {
	tree := sortedArrayToBST([]int{1, 3})
	ans := reflect.DeepEqual(BFS(tree), []int{3, 1})
	if !ans {
		t.Error("Wrong Answer!")
	}
}
