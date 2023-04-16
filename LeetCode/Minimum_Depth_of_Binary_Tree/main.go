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
* min Depth
*
* args :
*   root *TreeNode
*
* return value:
*   int
*
*
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := minDepth(root.Left)
	r := minDepth(root.Right)

	if root.Left == nil {
		return r + 1
	}
	if root.Right == nil {
		return l + 1
	}
	if l > r {
		return r + 1
	}
	return l + 1
}

func main() {
	// simple test
	node := &TreeNode{1, &TreeNode{}, &TreeNode{
		Left: &TreeNode{Left: &TreeNode{}},
	}}
	if ans := minDepth(node); ans != 2 {
		log.Fatalf("Worn Answer! answer equal `2` not equal : %d", ans)
	}
	log.Println("Answer is True!")
}
