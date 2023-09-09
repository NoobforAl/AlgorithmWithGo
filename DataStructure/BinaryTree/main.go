package main

import (
	"fmt"
	"math/rand"

	"github.com/kr/pretty"
)

type Tree struct {
	Left  *Tree
	Val   int
	Right *Tree
}

func (tr *Tree) Traverse() {
	var tra func(t *Tree)
	tra = func(t *Tree) {
		if t == nil {
			return
		}

		tra(t.Left)
		fmt.Print(t.Val, " ")
		tra(t.Right)
	}
	tra(tr)
}

func (tr *Tree) Insert(v int) *Tree {
	var ins func(t *Tree, v int) *Tree
	ins = func(t *Tree, v int) *Tree {
		if t == nil {
			return &Tree{Val: v}
		}

		if v == t.Val {
			return t
		}

		if v < t.Val {
			t.Left = ins(t.Left, v)
			return t
		}

		t.Right = ins(t.Right, v)
		return t
	}

	return ins(tr, v)
}

func CreateNewTree(n int) *Tree {
	var t *Tree
	for i := 1; i < n*2; i++ {
		tmp := rand.Intn(n * 2)
		t = t.Insert(tmp)
	}
	return t
}

func main() {
	t := CreateNewTree(10)
	fmt.Println("Print Create Value")
	t.Traverse()

	t = t.Insert(-3)
	t = t.Insert(-10)

	fmt.Println("\nAfter Appends Value")
	t.Traverse()
	fmt.Println()

	fmt.Printf("%# v\n", pretty.Formatter(t))
}
