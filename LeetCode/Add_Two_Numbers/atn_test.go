package main

import "testing"

func TestCase1(t *testing.T) {
	l1 := &ListNode{Val: 9}
	l2 := &ListNode{Val: 9}

	t1, t2 := l1, l2

	for _, v := range []int{9, 9, 9, 9, 9, 9} {
		t1 = insert(t1, v)
	}
	for _, v := range []int{9, 9, 9} {
		t2 = insert(t2, v)
	}

	ans := addTwoNumbers(l1, l2)

	trueAns, i := []int{8, 9, 9, 9, 0, 0, 0, 1}, 0

	for ; ans != nil; ans = ans.Next {
		if ans.Val != trueAns[i] {
			t.Error("Wrong Answer!")
			return
		}
		i++
		if i >= len(trueAns) {
			break
		}
	}

	if ans != nil && ans.Next != nil {
		t.Error("Wrong Answer!")
		return
	}
}

func TestCase2(t *testing.T) {
	l1 := &ListNode{Val: 0}
	l2 := &ListNode{Val: 0}

	ans := addTwoNumbers(l1, l2)

	trueAns, i := []int{0}, 0

	for ; ans != nil; ans = ans.Next {
		if ans.Val != trueAns[i] {
			t.Error("Wrong Answer!")
			return
		}
		i++
		if i >= len(trueAns) {
			break
		}
	}

	if ans != nil && ans.Next != nil {
		t.Error("Wrong Answer!")
		return
	}
}
