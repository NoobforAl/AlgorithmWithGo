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
* is Symmetric
*
* args :
*   root *TreeNode
*
*
* return value:
*   bool
*
*
 */
func isSymmetric(root *TreeNode) bool {
	var f func(rootL, rootR *TreeNode) bool
	f = func(rootL, rootR *TreeNode) bool {
		if rootL == nil && rootR == nil {
			return true
		}
		if rootL == nil || rootR == nil {
			return false
		}
		return f(rootL.Left, rootR.Right) &&
			f(rootL.Right, rootR.Left) &&
			rootR.Val == rootL.Val
	}
	return f(root, root)
}

func main() {
	// simple test
	if ans := isSymmetric(&TreeNode{}); !ans {
		log.Fatalf("Worn Answer! answer equal `True` not equal : %v", ans)
	}
	log.Println("Answer is True!")
}
