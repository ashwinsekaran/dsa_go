package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// LeetCode 234 - Palindrome Linked List
//
// Problem:
//   Given the head of a singly linked list, determine if it reads the
//   same forwards and backwards.
//
// Example:
//   Input:  1 -> 2 -> 2 -> 1 -> nil
//   Output: true
//
//   Input:  1 -> 2 -> 3 -> 4 -> nil
//   Output: false
//
//   Explanation:
//   Find the middle with slow/fast pointers, reverse the second half
//   in place, then walk the first half and reversed second half
//   together comparing values.
//
// Pseudo code:
//   slow, fast = head, head
//   while fast != nil && fast.Next != nil:
//     slow = slow.Next; fast = fast.Next.Next
//   reverse the list starting at slow -> prev
//   walk head and prev together, comparing values; mismatch -> false

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// find middle node (slow lands at the start of the second half)
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	// reverse the second half
	var prev *ListNode = nil
	for slow != nil {
		next := slow.Next
		slow.Next = prev
		prev = slow
		slow = next
	}

	// compare first half with reversed second half
	first, second := head, prev
	for second != nil {
		if first.Val != second.Val {
			return false
		}
		first = first.Next
		second = second.Next
	}
	return true
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

func main() {
	cases := [][]int{
		{},              // empty        -> true
		{1},             // single       -> true
		{1, 2},          // two, diff    -> false
		{1, 1},          // two, same    -> true
		{1, 2, 1},       // odd palin    -> true
		{1, 2, 2, 1},    // even palin   -> true
		{1, 2, 3, 4},    // not palin    -> false
		{1, 2, 3, 2, 1}, // odd palin    -> true
	}

	for _, c := range cases {
		head := build(c)
		fmt.Printf("%-18v -> %v\n", c, isPalindrome(head))
	}
}
