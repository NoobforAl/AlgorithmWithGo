package main

import (
	"reflect"
	"testing"

	p "github.com/kr/pretty"
)

func TestCase1(t *testing.T) {
	head := &ListNode{1, &ListNode{2, &ListNode{2, nil}}}
	answer := &ListNode{1, &ListNode{2, nil}}
	ans := deleteDuplicates(head)

	var wrong bool = true
	t1, t2 := answer, ans
	for t1 != nil && t2 != nil {
		if t1.Val != t2.Val {
			wrong = false
			break
		}
		t1, t2 = t1.Next, t2.Next
	}

	if !wrong || (t1 != t2) {
		t.Fatalf("Worn Answer! answer equal: %# v\n not equal: %# v\n", p.Formatter(answer), p.Formatter(ans))
	}
	t.Log("Answer is True!")
}

func TestCase2(t *testing.T) {
	head := &ListNode{0, &ListNode{}}
	answer := &ListNode{}

	if ans := deleteDuplicates(head); !reflect.DeepEqual(ans, answer) {
		t.Fatalf("Worn Answer! answer equal: %# v\n not equal: %# v\n", p.Formatter(answer), p.Formatter(ans))
	}
	t.Log("Answer is True!")
}
