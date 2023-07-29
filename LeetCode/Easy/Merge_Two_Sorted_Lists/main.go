package main

import "fmt"

// listNode to be like This!
// listNode{ 1, listNode{ 2, nil } }
// node1 -> node2 -> nil
type ListNode struct {
	// value
	Val int

	// next node then next
	// node is ListNode ( self type )
	Next *ListNode
}

/*
*
*
* check two list is not empty
*
* args :
*	list1 , list2 *ListNode
*
* return value:
*	*ListNode
* 	error
*
*
 */
func checkTwoLists(list1, list2 *ListNode) (*ListNode, error) {
	if list1 == nil && list2 == nil {
		return nil, fmt.Errorf("two list is empty")
	}

	if list1 == nil {
		return list2, fmt.Errorf("list1 is empty")
	}

	if list2 == nil {
		return list1, fmt.Errorf("list2 is empty")
	}

	return nil, nil
}

/*
*
*
* this func for insert values
* to list
*
* args :
*	val int
*	last *ListNode
*
* return value:
*	*ListNode
*
*
 */
func insert(val int, last *ListNode) *ListNode {
	temp := new(ListNode)
	temp.Val = val
	temp.Next = nil
	last.Next = temp
	return last.Next
}

/*
*
*
* merge two list to one list
*
* args :
*	list1 , list2, temp *ListNode
*
*
 */
func merge(list1, list2, temp *ListNode) {
	// this continue if one list become nil!
	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			temp = insert(list2.Val, temp)
			list2 = list2.Next
		} else {
			temp = insert(list1.Val, temp)
			list1 = list1.Next
		}
	}

	// one loop in bottom just run!
	// if list1 have value in this loop
	// insert to temp
	for ; list1 != nil; list1 = list1.Next {
		temp = insert(list1.Val, temp)
	}

	// if list2 have value in this loop
	// insert to temp
	for ; list2 != nil; list2 = list2.Next {
		temp = insert(list2.Val, temp)
	}
}

/*
*
*
* merge two list to one list
*
* args :
*	list1 , list2 *ListNode
*
* return:
*	*ListNode
*
*
 */
func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	// if one ot two list empty this work
	// you can print error and see witch one list
	// is empty
	if node, err := checkTwoLists(list1, list2); err != nil {
		// fmt.Println(err.Error())
		return node
	}

	// make response value
	var resNode ListNode

	// make copy form resNode
	// with reference of value
	// for insert values
	temp := &resNode

	// copy of two list
	// because we don't have
	// change two list!
	l1, l2 := list1, list2

	// setup minimum firs value
	// of two list
	if l1.Val > l2.Val {
		temp.Val = l2.Val
		l2 = l2.Next
	} else {
		temp.Val = l1.Val
		l1 = l1.Next
	}

	// merge to list to temp
	merge(l1, l2, temp)

	// and return resNode
	return &resNode
}

func main() {
	// very simple test
	list1 := convertToList([]int{1, 2, 4})
	list2 := convertToList([]int{1, 3, 4})

	fmt.Println(
		checkSorted(
			mergeTwoLists(list1, list2),
			&[]int{1, 1, 2, 3, 4, 4},
		),
	)
}

// -------------------------------------------------------------------------------
// this func for make linked list form
// one array or slice
func convertToList(arr []int) *ListNode {
	head := new(ListNode)
	if len(arr) == 0 {
		return head
	}
	head.Val = arr[0]
	head.Next = nil

	last := head
	for i := 1; i < len(arr); i++ {
		last = insert(arr[i], last)
	}
	return head
}

// check sorted lists
func checkSorted(list *ListNode, ans *[]int) bool {
	i, t := 0, list
	for ; i < len(*ans) && t != nil; i, t = i+1, t.Next {
		if (*ans)[i] != t.Val {
			return false
		}
	}

	return t == nil && i == len(*ans)
}
