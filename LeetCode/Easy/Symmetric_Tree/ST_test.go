package main

import (
	"testing"
)

func TestCase1(t *testing.T) {
	l := &TreeNode{
		1,
		&TreeNode{Val: 3},
		&TreeNode{Val: 2},
	}

	r := &TreeNode{
		1,
		&TreeNode{Val: 2},
		&TreeNode{Val: 3},
	}

	node := &TreeNode{0, l, r}
	if ans := isSymmetric(node); !ans {
		t.Fatalf("Worn Answer! answer equal 'True' not equal : %v", ans)
	}
	t.Log("Answer is True!")
}

func TestCase2(t *testing.T) {
	l := &TreeNode{
		1,
		&TreeNode{Val: 2, Right: &TreeNode{Val: 2}},
		nil,
	}

	r := &TreeNode{
		1,
		&TreeNode{Val: 2, Right: &TreeNode{Val: 2}},
		nil,
	}

	node := &TreeNode{0, l, r}
	if ans := isSymmetric(node); ans {
		t.Fatalf("Worn Answer! answer equal 'False' not equal : %v", ans)
	}
	t.Log("Answer is True!")
}
