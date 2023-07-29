package main

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	tmp := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 3},
		},
	}
	answer := []int{1, 3, 2}
	if ans := inorderTraversal(tmp); !reflect.DeepEqual(ans, answer) {
		t.Fatalf("Worn Answer! answer equal '[1, 3, 2]' not equal : %v", ans)
	}
	t.Log("Answer is True!")
}

func TestCase2(t *testing.T) {
	tmp := &TreeNode{Val: 1}
	answer := []int{1}
	if ans := inorderTraversal(tmp); !reflect.DeepEqual(ans, answer) {
		t.Fatalf("Worn Answer! answer equal '[1]' not equal : %v", ans)
	}
	t.Log("Answer is True!")
}
