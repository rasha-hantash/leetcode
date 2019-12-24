package main

import (
	"fmt"
)

//ListNode for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	//a := []int{8,20,19,2,14,9}
	//b := []int{8,20,19,2,14,9}

	// fmt.Println(twoSum(b, 17))
	// l1 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{Val: 4}}}}
	// l2 := &ListNode{2, &ListNode{2, &ListNode{4, &ListNode{Val: 6}}}}

	// r3 := mergeTwoLists(l1, l2)
	// display(r3)
	c := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(c))
}

//displays the linked list
func display(list *ListNode) {
	for list != nil {
		fmt.Printf("%v ->", list.Val)
		list = list.Next
	}
	fmt.Println()
}
