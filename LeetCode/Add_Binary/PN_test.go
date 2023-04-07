package main

import (
	"testing"
)

func TestCase1(t *testing.T) {
	if ans := addBinary("1010", "1011"); ans != "10101" {
		t.Fatalf("Worn Answer! answer equal '10101' not equal : %v", ans)
	}
	t.Log("Answer is True!")
}

func TestCase2(t *testing.T) {
	if ans := addBinary("110111", "101"); ans != "111100" {
		t.Fatalf("Worn Answer! answer equal '111100' not equal : %v", ans)
	}
	t.Log("Answer is True!")
}
