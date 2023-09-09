package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type Stack struct {
	Node *Node
	Size uint
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(val int) bool {
	if s.Node == nil {
		s.Node = &Node{Val: val}
		s.Size++
		return true
	}

	tmp := &Node{val, nil}
	tmp.Next = s.Node
	s.Node = tmp
	s.Size++
	return true
}

func (s *Stack) Pop() (int, bool) {
	if s.Size == 0 {
		return 0, false
	}

	tmp := s.Node
	s.Node = s.Node.Next
	s.Size--
	return tmp.Val, true
}

func (s *Stack) Traverse() {
	t := s.Node
	if t == nil {
		fmt.Println("-> Empty list!")
	}

	for ; t != nil; t = t.Next {
		fmt.Printf("%d ->", t.Val)
	}

	fmt.Println()
}

func main() {
	stack := NewStack()

	stack.Push(000)
	stack.Push(100)
	stack.Push(200)
	stack.Push(300)
	stack.Push(400)
	stack.Push(500)

	// 500 ->400 ->300 ->200 ->100 ->0 ->
	stack.Traverse()

	// 500 400 300 200 100 0
	for i := 0; i < 6; i++ {
		v, _ := stack.Pop()
		fmt.Print(v, " ")
	}
	fmt.Println()

	//-> Empty list!
	stack.Traverse()
}
