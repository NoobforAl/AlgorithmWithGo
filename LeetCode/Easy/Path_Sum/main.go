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
* has Path Sum
*
* args :
*   root *TreeNode
*   targetSum int
*
*
* return value:
*   bool
*
*
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	targetSum -= root.Val
	if targetSum == 0 && root.Left == nil && root.Right == nil {
		return true
	}
	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
}

func main() {
	// simple test
	if ans := hasPathSum(&TreeNode{1, nil, nil}, 1); !ans {
		log.Fatalf("Worn Answer! answer equal `True` not equal : %v", ans)
	}
	log.Println("Answer is True!")
}
