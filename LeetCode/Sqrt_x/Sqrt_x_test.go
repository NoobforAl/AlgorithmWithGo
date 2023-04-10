package main

import (
	"testing"
)

func TestCase1(t *testing.T) {
	if ans := mySqrt(8); ans != 2 {
		t.Fatalf("Worn Answer! answer equal '2' not equal : %d", ans)
	}
	t.Log("Answer is True!")
}

func TestCase2(t *testing.T) {
	if ans := mySqrt(4); ans != 2 {
		t.Fatalf("Worn Answer! answer equal '2' not equal : %d", ans)
	}
	t.Log("Answer is True!")
}
