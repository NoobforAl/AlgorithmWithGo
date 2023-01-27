package main

import (
	"fmt"
	"reflect"
	"testing"
)

/// [-1 0 1 0 1]

func TestCase1(t *testing.T) {
	example := [][]uint{
		{0, 2, 0, 6, 0},
		{2, 0, 3, 8, 5},
		{0, 3, 0, 0, 7},
		{6, 8, 0, 0, 9},
		{0, 5, 7, 9, 0},
	}

	answer := []int{-1, 0, 1, 0, 1}

	var ans []int = *prim(&example)

	if !reflect.DeepEqual(ans, answer) {
		fmt.Println(ans, answer)
		t.Error("Something was Wrong!")
	}
}
