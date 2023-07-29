package main

import (
	"log"
	"math"
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
* is Balanced
*
* args :
*   root *TreeNode
*
* return value:
*   bool
*
*
 */
func isBalanced(root *TreeNode) bool {
	var f func(r *TreeNode) int
	f = func(r *TreeNode) int {
		if r == nil {
			return 0
		}

		lh := f(r.Left)
		if lh == -1 {
			return -1
		}

		rh := f(r.Right)
		if rh == -1 {
			return -1
		}

		if int(math.Abs(float64(lh-rh))) > 1 {
			return -1
		}

		if rh > lh {
			return rh + 1
		}
		return lh + 1
	}
	return f(root) > -1
}

func main() {
	// simple test
	if ans := isBalanced(nil); !ans {
		log.Fatalf("Worn Answer! answer equal `True` not equal : %v", ans)
	}
	log.Println("Answer is True!")
}
