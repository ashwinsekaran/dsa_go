package main

import "sort"

// LeetCode 252 - Meeting Rooms
//
// Problem:
//   Given an array of meeting time intervals, determine if a person
//   could attend all meetings without any overlap.
//
// Example:
//   Input:  intervals = [[10, 12], [6, 9], [13, 15]]
//   Output: true
//
//   Explanation:
//   Sorted by start time: [6, 9], [10, 12], [13, 15]. Each meeting
//   starts at or after the previous one ends, so there's no conflict.
//
// Pseudo code:
//   sort intervals by start time
//   for i in 1..n-1:
//     if intervals[i].start < intervals[i-1].end: return false
//   return true

func main() {
	// intervals := [][]int{{1, 5}, {3, 9}, {6, 8}}
	intervals := [][]int{{10, 12}, {6, 9}, {13, 15}}
	result := canAttendMeetings(intervals)
	println(result)
}

func canAttendMeetings(intervals [][]int) bool {
	if len(intervals) == 0 {
		return true
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			return false
		}
	}

	return true
}
