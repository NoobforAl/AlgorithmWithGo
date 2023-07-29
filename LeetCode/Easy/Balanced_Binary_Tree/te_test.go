package main

import (
	"testing"
)

func TestCase1(t *testing.T) {
	value := &TreeNode{
		Val:   3,
		Right: &TreeNode{Val: 3},
		Left: &TreeNode{
			Val:   5,
			Right: &TreeNode{Val: 3},
			Left:  &TreeNode{Val: 4},
		},
	}
	if ans := isBalanced(value); !ans {
		t.Fatalf("Worn Answer! answer equal 'True' not equal : %v", ans)
	}
	t.Log("Answer is True!")
}

func TestCase2(t *testing.T) {
	value := &TreeNode{
		Val:   3,
		Right: &TreeNode{Val: 3},
		Left: &TreeNode{
			Val:   5,
			Right: &TreeNode{Val: 3},
			Left: &TreeNode{
				Val:   4,
				Right: &TreeNode{Val: 7},
				Left:  &TreeNode{Val: 8},
			},
		},
	}
	if ans := isBalanced(value); ans {
		t.Fatalf("Worn Answer! answer equal 'False' not equal : %v", ans)
	}
	t.Log("Answer is True!")
}
