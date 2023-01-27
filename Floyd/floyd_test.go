package main

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	example := [][]uint{
		{0, 5, INF, 10},
		{INF, 0, 3, INF},
		{INF, INF, 0, 1},
		{INF, INF, INF, 0},
	}

	answer := [][]uint{
		{0, 5, 8, 9},
		{6, 0, 3, 4},
		{3, 4, 0, 1},
		{3, 4, 7, 0},
	}

	floyd(&example)

	if !reflect.DeepEqual(answer, example) {
		t.Error("Something was wrong!")
	}
}

func TestCase2(t *testing.T) {
	example := [][]uint{
		{0, 8, 10, 15, 20},
		{15, 0, INF, 10, 8},
		{4, INF, 0, 8, INF},
		{INF, 6, 6, 0, INF},
		{15, 3, INF, INF, 0},
	}

	answer := [][]uint{
		{0, 8, 10, 15, 16},
		{15, 0, 16, 10, 8},
		{4, 12, 0, 8, 20},
		{10, 6, 6, 0, 14},
		{15, 3, 19, 13, 0},
	}

	floyd(&example)

	if !reflect.DeepEqual(answer, example) {
		t.Error("Something was wrong!")
	}
}
