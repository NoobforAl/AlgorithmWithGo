package main

import (
	"log"
)

// list Node
type ListNode struct {
	// value type int
	Val int

	// Next Node *ListNode
	Next *ListNode
}

/*
*
*
* insert new node to last node
*
* args :
*   last *ListNode
*   val int
*
* return value:
*	*ListNode
*
*
 */
func insert(last *ListNode, val int) *ListNode {
	temp := new(ListNode)
	temp.Val = val
	temp.Next = nil
	last.Next = temp
	return last.Next
}

/*
*
*
* sum two list numbers
*
* args :
*   l1, l2 *ListNode
*
* return value:
*   *ListNode
*
*
 */
func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	// make response node
	// make c ( carry )
	// sum value two int
	// first setup sum two value l1.Val + l2.Val
	res, c, s := &ListNode{Val: l1.Val + l2.Val}, 0, 0

	// make temp1 and temp2
	// we get first number from two list
	// because we need next node
	t1, t2 := l1.Next, l2.Next

	// if res.Val ( l1.Val + l2.Val ) > 9
	// we have carry
	if res.Val > 9 {
		c = res.Val / 10
		res.Val = res.Val % 10
	}

	// make temp for response value
	temp := res

	// loop to t1 and t2
	// if t1 or t2 equal nil
	// loop break
	for ; t1 != nil && t2 != nil; t1, t2 = t1.Next, t2.Next {

		// sum tow value t1.Val + t2.Val + c
		s = t1.Val + t2.Val + c
		// make carry zero
		c = 0
		// check s > 0 orn ot
		if s > 9 {
			// change c and s
			// make s less 10 ( s % 10)
			// setup new carry s / 10
			c, s = s/10, s%10
		}

		// and insert
		temp = insert(temp, s)
	}

	// this loop for insert
	// value form t1 != nil
	for ; t1 != nil; t1 = t1.Next {
		s = t1.Val + c
		c = 0
		if s > 9 {
			c, s = s/10, s%10
		}
		temp = insert(temp, s)
	}

	// this loop for insert
	// value form t2 != nil
	for ; t2 != nil; t2 = t2.Next {
		s = t2.Val + c
		c = 0
		if s > 9 {
			c, s = s/10, s%10
		}
		temp = insert(temp, s)
	}

	// if c != 0 insert New value to list
	if c != 0 {
		temp = insert(temp, c)
	}

	return res
}

func main() {
	// very simple test
	l1 := &ListNode{Val: 2}
	l2 := &ListNode{Val: 5}

	t1, t2 := l1, l2

	for _, v := range []int{4, 3} {
		t1 = insert(t1, v)
	}
	for _, v := range []int{6, 4} {
		t2 = insert(t2, v)
	}

	ans := addTwoNumbers(l1, l2)

	trueAns, i := []int{7, 0, 8}, 0

	for ; ans != nil; ans = ans.Next {
		if ans.Val != trueAns[i] {
			log.Fatal("Wrong Answer!")
		}
		i++
		if i >= len(trueAns) {
			break
		}
	}

	if ans != nil && ans.Next != nil {
		log.Fatal("Wrong Answer!")
	}

	log.Println("True Answer!")
}
