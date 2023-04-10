package main

import (
	"testing"
)

func TestCase1(t *testing.T) {
	if ans := hasPathSum(&TreeNode{1, &TreeNode{}, nil}, 1); !ans {
		t.Fatalf("Worn Answer! answer equal 'True' not equal : %v", ans)
	}
	t.Log("Answer is True!")
}

func TestCase2(t *testing.T) {
	if ans := hasPathSum(&TreeNode{3, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}, 5); !ans {
		t.Fatalf("Worn Answer! answer equal 'True' not equal : %v", ans)
	}
	t.Log("Answer is True!")
}
