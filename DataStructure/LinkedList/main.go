package main

import "fmt"

type LinkedList struct {
	Val  int
	Next *LinkedList
}

func (l *LinkedList) AddNode(val int) int {
	var rev func(t *LinkedList, val int) int
	rev = func(t *LinkedList, val int) int {
		if t == nil {
			t = &LinkedList{Val: val}
			return val
		}

		if val == t.Val {
			fmt.Println("Node already exists: ", val)
			return -1
		}

		if t.Next == nil {
			t.Next = &LinkedList{Val: val}
			return -2
		}

		return rev(t.Next, val)
	}
	return rev(l, val)
}

func (l *LinkedList) Traverse() {
	t := l
	if t == nil {
		fmt.Println("-> Empty list!")
	}

	for t != nil {
		fmt.Printf("%d ->", t.Val)
		t = t.Next
	}
	fmt.Println()
}

func (l *LinkedList) LookupNode(val int) bool {
	var rev func(t *LinkedList, val int) bool
	rev = func(t *LinkedList, val int) bool {
		if val == t.Val {
			return true
		}

		if t.Next == nil {
			return false
		}

		return rev(t.Next, val)
	}
	return rev(l, val)
}

func (l *LinkedList) Size() (len int) {
	for t := l; t != nil; t = t.Next {
		len++
	}
	return
}

func main() {
	l := new(LinkedList)

	l.AddNode(10)
	l.AddNode(5)
	l.AddNode(45)
	l.AddNode(5)
	l.AddNode(5)
	l.Traverse()

	if l.LookupNode(45) {
		fmt.Println("Node exists!")
	} else {
		fmt.Println("Node does not exist!")
	}

	fmt.Println("Size: ", l.Size())
}
