package main

// LeetCode 1423 - Maximum Points You Can Obtain from Cards
//
// Problem:
//   There are several cards arranged in a row. Each card has a point value.
//   In one step you can take one card from the beginning or end of the row.
//   You must take exactly k cards. Find the maximum score you can obtain.
//
// Example:
//   Input:  cards = [100, 40, 17, 9, 73, 75], k = 3
//   Output: 248
//
//   Explanation:
//   Taking the k cards from the ends is equivalent to leaving behind a
//   contiguous window of size len(cards)-k in the middle. Minimizing that
//   leftover window's sum maximizes the picked score: total - min(window).
//   Here the leftover [17, 9] (or similar) minimizes to total-248's complement.
//
// Pseudo code:
//   total = sum(cards)
//   windowSum = sum of first (len(cards)-k) cards   # the complement window
//   slide window of size len(cards)-k across cards
//   maxSum = max(maxSum, total - windowSum)         # picks = total - min complement

func main() {
	nums := []int{100, 40, 17, 9, 73, 75}
	k := 3

	res := maxPoints(nums, k)
	println(res) // 248
}

func maxPoints(cards []int, k int) int {
	maxSum := 0
	windowSum := 0
	start := 0
	total := 0

	if len(cards) > 0 {
		for _, num := range cards {
			total += num
		}
	}

	if k == len(cards) {
		return total
	}

	for end := 0; end < len(cards); end++ {
		windowSum += cards[end]

		if end-start+1 == len(cards)-k {
			if total-windowSum > maxSum {
				maxSum = total - windowSum
			}
			windowSum -= cards[start]
			start++
		}
	}
	return maxSum
}
