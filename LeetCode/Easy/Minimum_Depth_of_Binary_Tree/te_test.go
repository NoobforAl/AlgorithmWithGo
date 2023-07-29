package main

import (
	"testing"
)

func TestCase1(t *testing.T) {
	node := &TreeNode{3, &TreeNode{Val: 9}, &TreeNode{10, &TreeNode{}, &TreeNode{}}}
	if ans := minDepth(node); ans != 2 {
		t.Fatalf("Worn Answer! answer equal '2' not equal : %v", ans)
	}
	t.Log("Answer is True!")
}

func TestCase2(t *testing.T) {
	node := &TreeNode{1, nil,
		&TreeNode{2, nil,
			&TreeNode{3, nil,
				&TreeNode{4, nil,
					&TreeNode{5, nil,
						&TreeNode{6, nil, nil},
					},
				},
			},
		},
	}
	if ans := minDepth(node); ans != 6 {
		t.Fatalf("Worn Answer! answer equal '6' not equal : %v", ans)
	}
	t.Log("Answer is True!")
}
