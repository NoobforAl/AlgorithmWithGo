package main

import (
	"testing"
)

func TestCase1(t *testing.T) {
	node := &TreeNode{}
	if ans := maxDepth(node); ans != 1 {
		t.Fatalf("Worn Answer! answer equal '1' not equal : %v", ans)
	}
	t.Log("Answer is True!")
}

func TestCase2(t *testing.T) {
	node := &TreeNode{
		Left: &TreeNode{
			Left: &TreeNode{
				Left: &TreeNode{
					Left: &TreeNode{
						Left: &TreeNode{}}}}}}
	if ans := maxDepth(node); ans != 6 {
		t.Fatalf("Worn Answer! answer equal '6' not equal : %v", ans)
	}
	t.Log("Answer is True!")
}
