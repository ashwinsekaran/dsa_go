package main

// LeetCode 2461 - Maximum Sum of Distinct Subarrays With Length K
//
// Problem:
//   Given an integer array nums and an integer k, find the maximum subarray
//   sum of all subarrays of length k that contain only distinct elements.
//   If no such subarray exists, return 0.
//
// Example:
//   Input:  nums = [3, 2, 2, 3, 4, 6, 7, 7, -1], k = 4
//   Output: 20
//
//   Explanation:
//   The window [3, 4, 6, 7] (indices 3-6) has 4 distinct elements and
//   sums to 20, which is the maximum among all valid windows.
//
// Pseudo code:
//   windowSum += nums[end]; freqMap[nums[end]]++
//   if window size == k:
//     if len(freqMap) == k: maxDistinctSum = max(maxDistinctSum, windowSum)
//     windowSum -= nums[start]; decrement/delete freqMap[nums[start]]; start++

func main() {
	nums := []int{3, 2, 2, 3, 4, 6, 7, 7, -1}
	k := 4
	res := maxDistinctSumSubArray(nums, k)
	println(res) // 20
}

func maxDistinctSumSubArray(nums []int, k int) int {
	start := 0
	windowSum := 0
	maxDistinctSum := 0
	freqMap := make(map[int]int)

	for end := 0; end < len(nums); end++ {
		windowSum += nums[end]
		freqMap[nums[end]]++

		if end-start+1 == k {
			if len(freqMap) == k {
				if windowSum > maxDistinctSum {
					maxDistinctSum = windowSum
				}
			}
			windowSum -= nums[start]
			freqMap[nums[start]]--
			if freqMap[nums[start]] == 0 {
				delete(freqMap, nums[start])
			}
			start++
		}

	}
	return maxDistinctSum
}
