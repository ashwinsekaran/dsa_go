package main

import "fmt"

// LeetCode 11 - Container With Most Water
//
// Problem:
//   Given n non-negative integers representing wall heights, find two
//   lines that together with the x-axis form a container holding the
//   most water.
//
// Example:
//   Input:  heights = [3, 4, 1, 2, 2, 4, 1, 3, 2]
//   Output: 21
//
//   Explanation:
//   Lines at index 0 (height 3) and index 7 (height 3).
//   Width = 7 - 0 = 7, effective height = min(3,3) = 3.
//   Area = 7 * 3 = 21.
//   Strategy: always move the pointer at the shorter line inward, since
//   moving the taller one can only keep the height the same or shrink it.
//
// Pseudo code:
//   left, right = 0, len(heights)-1
//   area = min(heights[left], heights[right]) * (right - left)
//   advance the pointer at the shorter height
//   return max area seen

func main() {
	heights := []int{3, 4, 1, 2, 2, 4, 1, 3, 2}

	res := maxArea(heights)

	fmt.Println(res)
}

func maxArea(heights []int) int {
	currentMax := 0
	if len(heights) > 0 {
		left, right := 0, len(heights)-1

		for left < right {
			width := right - left
			height := min(heights[left], heights[right])
			currentArea := width * height

			if currentArea > currentMax {
				currentMax = currentArea
			}
			if heights[left] < heights[right] {
				left++
			} else {
				right--
			}

		}

	}
	return currentMax
}
