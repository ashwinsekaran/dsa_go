package main

import (
	"fmt"
	"sort"
)

// Pair with Target Sum (Two Pointers)
//
// Problem:
//   Given an array of integers and a target sum, determine whether any
//   pair of elements adds up to the target.
//
// Example:
//   Input:  nums = [1, 1, 3, 4, 6], target = 2
//   Output: true
//
//   Explanation:
//   After sorting, [1, 1, 3, 4, 6]. The pair at the two 1s (indices 0, 1)
//   sums to 2, matching the target.
//   Strategy: walk pointers inward from both ends — if the sum is too
//   small, move left forward; if too large, move right backward.
//
// Pseudo code:
//   sort(nums)
//   left, right = 0, len(nums)-1
//   while left < right:
//     sum = nums[left] + nums[right]
//     if sum == target: return true
//     if sum > target: right--
//     else: left++
//   return false

func main() {
	//nums := []int{1, 3, 4, 6, 8, 10, 13}
	//target := 13

	nums := []int{1, 1, 3, 4, 6}
	target := 2

	res := findTargetSumOfAPair(nums, target)
	fmt.Println(res)
}

func findTargetSumOfAPair(nums []int, target int) bool {
	sort.Ints(nums)
	left, right := 0, len(nums)-1

	for left < right {
		currentSum := nums[left] + nums[right]
		if currentSum == target {
			return true
		}
		if currentSum > target {
			right--
		} else {
			left++
		}
	}
	return false
}
