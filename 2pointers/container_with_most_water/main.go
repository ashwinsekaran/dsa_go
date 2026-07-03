package main

import "fmt"

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
