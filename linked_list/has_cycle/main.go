package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// LeetCode 141 - Linked List Cycle
//
// Problem:
//   Given the head of a linked list, determine if it has a cycle —
//   some node's Next eventually points back to a previously visited node.
//
// Example:
//   Input:  1 -> 2 -> 3 -> 4 -> nil                (straight)
//   Output: false
//
//   Input:  1 -> 2 -> 3 -> 4 -> (back to 2)         (looped at index 1)
//   Output: true
//
//   Explanation:
//   Track every visited node in a set. If a node is seen twice, the
//   list loops; reaching nil means it terminates normally.
//
// Pseudo code:
//   visited = {}
//   for node := head; node != nil; node = node.Next:
//     if visited[node]: return true
//     visited[node] = true
//   return false

func hasCycle(head *ListNode) bool {
	visitedNodes := make(map[*ListNode]bool)
	currentNode := head
	for currentNode != nil {
		if visitedNodes[currentNode] {
			return true // cycle detected
		}
		visitedNodes[currentNode] = true
		currentNode = currentNode.Next
	}
	return false
}

// helper: build a straight linked list from a slice
func build(vals []int) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for _, v := range vals {
		tail.Next = &ListNode{Val: v}
		tail = tail.Next
	}
	return dummy.Next
}

// helper: make the last node point back to the node at position `pos`
// (0-indexed). pos = -1 means no cycle.
func makeCycle(head *ListNode, pos int) *ListNode {
	if head == nil {
		return nil
	}

	// find the node at index `pos`
	var target *ListNode
	if pos >= 0 {
		target = head
		for i := 0; i < pos; i++ {
			target = target.Next
		}
	}

	// walk to the last node
	tail := head
	for tail.Next != nil {
		tail = tail.Next
	}

	tail.Next = target // link last node back to target (or nil)
	return head
}

func main() {
	// no cycle
	straight := build([]int{1, 2, 3, 4})
	fmt.Println("straight list has cycle:", hasCycle(straight))

	// with a cycle: last node points back to index 1 (value 2)
	looped := build([]int{1, 2, 3, 4})
	looped = makeCycle(looped, 1)
	fmt.Println("looped list has cycle:  ", hasCycle(looped))
}
