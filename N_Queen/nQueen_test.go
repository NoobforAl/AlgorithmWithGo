package main

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	example := [][]int8{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	answer := [][]int8{
		{0, 1, 0, 0},
		{0, 0, 0, 1},
		{1, 0, 0, 0},
		{0, 0, 1, 0},
	}

	nQueen(&example, 0)

	if !reflect.DeepEqual(example, answer) {
		t.Error("Something was Wrong!")
	}
}

func TestCase2(t *testing.T) {
	example := [][]int8{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}

	answer := [][]int8{
		{1, 0, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 0, 1},
		{0, 1, 0, 0, 0},
		{0, 0, 0, 1, 0},
	}

	nQueen(&example, 0)
	if !reflect.DeepEqual(example, answer) {
		t.Error("Something was Wrong!")
	}
}
