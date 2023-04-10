package main

import (
	"reflect"
	"testing"

	pr "github.com/kr/pretty"
)

var m = Map{
	{0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}

func TestCase1(t *testing.T) {
	start := Node{Location: Point{2, 4}}
	end := Node{Location: Point{8, 6}}

	out := aStar(m, &start, &end)
	ans := []Point{
		{8, 6}, {7, 7},
		{6, 7}, {5, 8},
		{5, 9}, {6, 9},
		{7, 8}, {8, 7},
		{8, 8}, {9, 7},
		{9, 8}, {9, 9},
		{8, 9}, {7, 9},
		{6, 8}, {5, 7},
		{4, 6}, {3, 5},
		{2, 4}}

	if !reflect.DeepEqual(ans, out) {
		t.Fatalf("This Answer Not Equal:\n %# v \nYou Answer:\n %# v",
			pr.Formatter(out),
			pr.Formatter(ans))
	}
	t.Log("True Answer!")
}

func TestCase2(t *testing.T) {
	start := Node{Location: Point{2, 2}}
	end := Node{Location: Point{8, 8}}

	out := aStar(m, &start, &end)
	ans := []Point{
		{8, 8}, {8, 9},
		{9, 9}, {9, 8},
		{8, 7}, {7, 6},
		{6, 5}, {5, 4},
		{4, 3}, {3, 3},
		{2, 2}}

	if !reflect.DeepEqual(ans, out) {
		t.Fatalf("This Answer Not Equal:\n %# v \nYou Answer:\n %# v",
			pr.Formatter(out),
			pr.Formatter(ans))
	}
	t.Log("True Answer!")
}
