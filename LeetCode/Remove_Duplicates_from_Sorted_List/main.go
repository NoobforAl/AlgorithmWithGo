package main

import (
	"log"
)

// List Node
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
*
*
* delete Duplicates
*
* args :
*   head *ListNode*
*
* return value:
*   *ListNode
*
*
 */
func deleteDuplicates(head *ListNode) *ListNode {
	for tmp := head; tmp != nil && tmp.Next != nil; {
		if tmp.Val == tmp.Next.Val {
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}
	}
	return head
}

func main() {
	// simple test
	if ans := deleteDuplicates(nil); ans != nil {
		log.Fatalf("Worn Answer! answer equal `nil` not equal : %v", ans)
	}
	log.Println("Answer is True!")
}
