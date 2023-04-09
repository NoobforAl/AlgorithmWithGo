package main

import (
	"log"
)

// Tree Node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
*
*
* is Same Tree
*
* args :
*   p, q *TreeNode
*
*
* return value:
*   bool
*
*
 */
func isSameTree(p, q *TreeNode) bool {
	switch {
	case p == nil && q == nil:
		return true
	case p == nil || q == nil || p.Val != q.Val:
		return false
	default:
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
}

func main() {
	// simple test
	p := &TreeNode{1, nil, &TreeNode{1, nil, nil}}
	q := &TreeNode{1, &TreeNode{1, nil, nil}, nil}

	if ans := isSameTree(p, q); ans {
		log.Fatalf("Worn Answer! answer equal `False` not equal : %v", ans)
	}
	log.Println("Answer is True!")
}
