package main

import (
	"testing"
)

func TestCase1(t *testing.T) {
	p := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	q := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}

	if ans := isSameTree(p, q); !ans {
		t.Fatalf("Worn Answer! answer equal 'True' not equal : %v", ans)
	}
	t.Log("Answer is True!")
}

func TestCase2(t *testing.T) {
	p := &TreeNode{1, &TreeNode{3, nil, nil}, &TreeNode{2, nil, nil}}
	q := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}

	if ans := isSameTree(p, q); ans {
		t.Fatalf("Worn Answer! answer equal 'False' not equal : %v", ans)
	}
	t.Log("Answer is True!")
}
