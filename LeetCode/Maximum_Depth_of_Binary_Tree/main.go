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
* max Depth
*
* args :
*   root *TreeNode
*
* return value:
*   int
*
*
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	if r > l {
		return r + 1
	}
	return l + 1
}

func main() {
	// simple test
	if ans := maxDepth(nil); ans != 0 {
		log.Fatalf("Worn Answer! answer equal `0` not equal : %v", ans)
	}
	log.Println("Answer is True!")
}
