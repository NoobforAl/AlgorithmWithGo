package main

import "fmt"

type Node struct {
	Val      int
	Next     *Node
	Previous *Node
}

func (n *Node) AddNode(val int) int {
	var rev func(t *Node, val int) int
	rev = func(t *Node, val int) int {
		if t == nil {
			t = &Node{Val: val}
			return 0
		}

		if val == t.Val {
			fmt.Println("Node already exists: ", val)
			return -1
		}

		if t.Next == nil {
			tmp := t
			t.Next = &Node{Val: val, Previous: tmp}
			return -2
		}
		return rev(t.Next, val)
	}
	return rev(n, val)
}

func (n *Node) Traverse() {
	if n == nil {
		fmt.Println("-> Empty list!")
	}

	for t := n; t != nil; t = t.Next {
		fmt.Printf("%d ->", t.Val)
	}
	fmt.Println()
}

func (n *Node) Reverse() {
	if n == nil {
		fmt.Println("-> Empty list!")
	}

	tmp := n
	for ; tmp.Next != nil; tmp = tmp.Next {
	}

	for ; tmp.Previous != nil; tmp = tmp.Previous {
		fmt.Printf("%d ->", tmp.Val)
	}
	fmt.Printf("%d ->\n", tmp.Val)
}

func (n *Node) LookupNode(val int) bool {
	var rev func(t *Node, val int) bool
	rev = func(t *Node, val int) bool {
		if t == nil {
			return false
		}

		if val == t.Val {
			return true
		}

		if t.Next == nil {
			return false
		}
		return rev(t.Next, val)
	}
	return rev(n, val)
}

func (n *Node) Size() (len int) {
	for t := n; t != nil; t = t.Next {
		len++
	}
	return
}

func main() {
	n := new(Node)

	n.AddNode(10)
	n.AddNode(5)
	n.AddNode(45)

	// Node already exists:  5
	n.AddNode(5)

	// Node already exists:  5
	n.AddNode(5)

	// 0 ->10 ->5 ->45 ->
	n.Traverse()

	//Size: 4
	fmt.Println("Size:", n.Size())

	// 45 ->5 ->10 ->0 ->
	n.Reverse()
}
