package main

// Maximum Sum Subarray of Size K (classic sliding window problem)
//
// Problem:
//   Given an integer array nums and an integer k, find the maximum sum
//   among all contiguous subarrays of size exactly k.
//
// Example:
//   Input:  nums = [2, 1, 5, 1, 3, 2], k = 3
//   Output: 9
//
//   Explanation:
//   The window [5, 1, 3] (indices 2-4) sums to 9, the maximum among all
//   size-3 windows.
//
// Pseudo code:
//   windowSum += nums[end]
//   if window size == k:
//     maxSum = max(maxSum, windowSum)
//     windowSum -= nums[start]; start++

func main() {
	nums := []int{2, 1, 5, 1, 3, 2}
	k := 3

	res := maxSubArray(nums, k)
	println(res) // 9
}

func maxSubArray(nums []int, k int) int {
	maxSum := 0
	windowSum := 0
	start := 0

	for end := 0; end < len(nums); end++ {
		windowSum += nums[end]

		if end-start+1 == k {
			if windowSum > maxSum {
				maxSum = windowSum
			}

			windowSum -= nums[start]
			start++
		}
	}
	return maxSum
}
