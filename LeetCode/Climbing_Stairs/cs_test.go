package main

import (
	"testing"
)

func TestCase1(t *testing.T) {
	if ans := climbStairs(44); ans != 1134903170 {
		t.Fatalf("Worn Answer! answer equal '1134903170' not equal : %v", ans)
	}
	t.Log("Answer is True!")
}

func TestCase2(t *testing.T) {
	if ans := climbStairs(45); ans != 1836311903 {
		t.Fatalf("Worn Answer! answer equal 'True' not equal : %v", ans)
	}
	t.Log("Answer is True!")
}
