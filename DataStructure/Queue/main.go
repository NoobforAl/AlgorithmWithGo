package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type Queue struct {
	Node *Node
	Size uint
}

func (q *Queue) Push(val int) bool {
	if q.Node == nil {
		q.Node = &Node{Val: val}
		q.Size++
		return true
	}

	tmp := &Node{val, nil}
	tmp.Next = q.Node
	q.Node = tmp
	q.Size++
	return true
}

func (q *Queue) Pop() (int, bool) {
	if q.Size == 0 {
		return 0, false
	}

	var tmp *Node = q.Node
	if q.Size == 1 {
		q.Node = nil
		q.Size = 0
		return tmp.Val, true
	}

	for t := q.Node; t.Next != nil; t = t.Next {
		tmp = t
	}

	v := tmp.Next.Val
	tmp.Next = nil
	q.Size--
	return v, true
}

func (q *Queue) Traverse() {
	t := q.Node
	if t == nil {
		fmt.Println("-> Empty list!")
	}

	for ; t != nil; t = t.Next {
		fmt.Printf("%d ->", t.Val)
	}
	fmt.Println()
}

func main() {
	q := new(Queue)

	q.Push(00)
	q.Push(10)
	q.Push(20)
	q.Push(30)
	q.Push(40)
	q.Push(50)

	fmt.Println("Size:", q.Size)
	q.Traverse()

	for i := 0; i < 6; i++ {
		v, _ := q.Pop()
		fmt.Printf("-> %d ", v)
	}

	fmt.Println()
	fmt.Println("Size:", q.Size)
	q.Traverse()
}
