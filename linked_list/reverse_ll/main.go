package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// LeetCode 206 - Reverse Linked List
//
// Problem:
//   Given the head of a singly linked list, reverse the list and
//   return the new head.
//
// Example:
//   Input:  1 -> 2 -> 3 -> 4 -> 5 -> nil
//   Output: 5 -> 4 -> 3 -> 2 -> 1 -> nil
//
//   Explanation:
//   Walk the list once, re-pointing each node's Next to the previous
//   node instead of the next one. prev starts at nil (the new tail's
//   Next) and ends up at the old tail, which becomes the new head.
//
// Pseudo code:
//   prev = nil; current = head
//   while current != nil:
//     next = current.Next
//     current.Next = prev
//     prev = current
//     current = next
//   return prev

func reverse(head *ListNode) *ListNode {
	var prev *ListNode = nil
	current := head

	for current != nil {
		next := current.Next // save next node before overwriting
		current.Next = prev  // reverse the pointer
		prev = current       // advance prev
		current = next       // advance current
	}

	return prev // prev is the new head
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
func print(head *ListNode) {
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
	fmt.Println("nil")
}

func main() {
	head := build([]int{1, 2, 3, 4, 5})

	fmt.Print("original: ")
	print(head)

	reversed := reverse(head)

	fmt.Print("reversed: ")
	print(reversed)
}
