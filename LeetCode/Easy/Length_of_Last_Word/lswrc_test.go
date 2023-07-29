package main

import (
	"testing"
)

func TestCase1(t *testing.T) {
	if ans := lengthOfLastWord("   fly me   to   the moon  "); ans != 4 {
		t.Fatalf("Worn Answer! answer equal 4 not equal : %d", ans)
	}
	t.Log("Answer is True!")
}

func TestCase2(t *testing.T) {
	if ans := lengthOfLastWord("luffy is still joyboy"); ans != 6 {
		t.Fatalf("Worn Answer! answer equal 6 not equal : %d", ans)
	}
	t.Log("Answer is True!")
}
