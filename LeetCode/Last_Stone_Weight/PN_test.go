package main

import (
	"bytes"
	"testing"

	"github.com/kr/pretty"
)

func TestCase1(t *testing.T) {
	var diff bytes.Buffer

	pretty.Fdiff(&diff, lastStoneWeight([]int{2, 7, 4, 1, 8, 1}), 1)

	if diff.String() != "" {
		t.Fatalf("Worn Answer!\n %v", diff.String())
	}
	t.Log("Answer is True!")
}

func TestCase2(t *testing.T) {
	var diff bytes.Buffer

	pretty.Fdiff(&diff, lastStoneWeight([]int{1}), 1)

	if diff.String() != "" {
		t.Fatalf("Worn Answer!\n %v", diff.String())
	}
	t.Log("Answer is True!")
}
