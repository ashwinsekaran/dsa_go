package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// LeetCode 21 - Merge Two Sorted Lists
//
// Problem:
//   Given the heads of two sorted linked lists, merge them into one
//   sorted list and return its head.
//
// Example:
//   Input:  l1 = 1 -> 3 -> 5 -> 7, l2 = 2 -> 4 -> 6 -> 8
//   Output: 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8
//
//   Explanation:
//   Repeatedly take the smaller of the two current nodes and splice it
//   onto the result, advancing that list's pointer. Once one list is
//   exhausted, attach the remainder of the other directly.
//
// Pseudo code:
//   pick smaller head as start; advance that list
//   while both lists non-empty:
//     attach smaller of l1.Val, l2.Val; advance that list
//   attach whichever list remains
//   return start

func mergeLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var head *ListNode
	if l1.Val < l2.Val {
		head = l1
		l1 = l1.Next
	} else {
		head = l2
		l2 = l2.Next
	}

	current := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}

	if l1 != nil {
		current.Next = l1
	} else {
		current.Next = l2
	}

	return head
}

// helper: build a linked list from a slice
func build(vals []int) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for _, v := range vals {
		tail.Next = &ListNode{Val: v}
		tail = tail.Next
	}
	return dummy.Next
}

// helper: print a linked list as 1 -> 2 -> 3 -> nil
func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
	fmt.Println("nil")
}

func main() {
	a := build([]int{1, 3, 5, 7})
	b := build([]int{2, 4, 6, 8})

	fmt.Print("list a:  ")
	printList(a)
	fmt.Print("list b:  ")
	printList(b)

	merged := mergeLists(a, b)

	fmt.Print("merged:  ")
	printList(merged)
}
