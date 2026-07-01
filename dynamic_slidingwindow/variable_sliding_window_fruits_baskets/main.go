package main

// LeetCode 904 - Fruit Into Baskets
//
// Problem:
//   You are visiting a row of fruit trees, fruits[i] is the type of fruit
//   on the ith tree. You have exactly 2 baskets, each can hold only one
//   type of fruit (unlimited amount). Find the length of the longest
//   contiguous subarray containing at most 2 distinct fruit types.
//
// Example:
//   Input:  fruits = [3, 3, 2, 1, 2, 1, 0]
//   Output: 4
//
//   Explanation:
//   The window [2, 1, 2, 1] (indices 2-5) has only 2 distinct types.
//   Whenever a 3rd type enters the window, shrink from the left by one
//   until that type's count drops out of the basket map.
//
// Pseudo code:
//   basket[fruits[end]]++
//   if len(basket) > 2: basket[fruits[start]]--; delete if zero; start++
//   maxFruit = max(maxFruit, end-start+1)

func main() {
	fruits := []int{3, 3, 2, 1, 2, 1, 0}

	res := maxFruits(fruits)
	println(res) // 4
}

func maxFruits(fruits []int) int {
	maxFruit := 0
	start := 0
	basket := make(map[int]int)

	for end := 0; end < len(fruits); end++ {
		basket[fruits[end]]++

		if len(basket) > 2 {
			basket[fruits[start]]--
			if basket[fruits[start]] == 0 {
				delete(basket, fruits[start])
			}
			start++
		}
		if end-start+1 > maxFruit {
			maxFruit = end - start + 1
		}
	}
	return maxFruit
}
