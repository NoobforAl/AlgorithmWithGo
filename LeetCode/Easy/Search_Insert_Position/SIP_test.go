package main

import (
	"testing"
)

func TestCase1(t *testing.T) {
	if ans := searchInsert([]int{1, 3, 5, 6}, 2); ans != 1 {
		t.Fatalf("Worn Answer! answer equal 1 not equal : %d", ans)
	}
	t.Log("Answer is True!")
}

func TestCase2(t *testing.T) {
	if ans := searchInsert([]int{1, 3, 5, 6}, 7); ans != 4 {
		t.Fatalf("Worn Answer! answer equal 4 not equal : %d", ans)
	}
	t.Log("Answer is True!")
}
