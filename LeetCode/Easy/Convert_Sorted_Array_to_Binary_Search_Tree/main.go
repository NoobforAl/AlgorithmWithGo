package main

import (
	"fmt"
	"reflect"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
*
*
* Sort array to BST
*
* args :
*   nums slices []int
*
* return value:
*   *TreeNode
*
*
 */
func sortedArrayToBST(nums []int) *TreeNode {
	// first check size nums is not equal zero
	if size := len(nums); size > 0 {

		// and return new node!
		// but this func work recursive
		return &TreeNode{

			// set mid value for value
			// new node
			Val: nums[size/2],

			// check left nums and call again func
			Left: sortedArrayToBST(nums[:size/2]),

			// check left nums and call again func
			Right: sortedArrayToBST(nums[size/2+1:]),
		}
	}

	// length nums is zero return nil
	return nil
}

func main() {
	// very simple test!
	t := sortedArrayToBST([]int{-10, -3, 0, 5, 9})
	ans := reflect.DeepEqual(BFS(t), []int{0, -3, 9, -10, 5})
	fmt.Println(ans)
}

/*
*
*
* bottom func for check
* answer is true or not!
*
* BFS Algorithm!
*
*
 */

func BFS(tree *TreeNode) []int {
	var queue []*TreeNode
	var result []int

	queue = append(queue, tree)
	return BFSUtil(queue, result)
}

func BFSUtil(queue []*TreeNode, res []int) []int {
	if len(queue) == 0 {
		return res
	}

	res = append(res, queue[0].Val)
	if queue[0].Left != nil {
		queue = append(queue, queue[0].Left)
	}

	if queue[0].Right != nil {
		queue = append(queue, queue[0].Right)
	}

	return BFSUtil(queue[1:], res)
}
