package main

import (
	"log"
	"reflect"
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
* inOrder Traversal
*
* args :
*   root *TreeNode
*
* return value:
*   []int slice
*
*
 */
func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	res = append(res, inorderTraversal(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)
	return res
}

func main() {
	// simple test
	if ans := inorderTraversal(nil); !reflect.DeepEqual(ans, []int{}) {
		log.Fatalf("Worn Answer! answer equal `[]` not equal : %v", ans)
	}
	log.Println("Answer is True!")
}

// Without Recursive
// We use Stack
func S_inOrderTraversal(root *TreeNode) []int {
	s := []*TreeNode{}
	res := []int{}
	tmp := root

	for {
		if len(s) == 0 && tmp == nil {
			break
		}

		for ; tmp != nil; tmp = tmp.Left {
			s = append(s, tmp)
		}

		tmp = s[len(s)-1]
		s = s[:len(s)-1]

		res = append(res, tmp.Val)
		tmp = tmp.Right
	}
	return res
}
