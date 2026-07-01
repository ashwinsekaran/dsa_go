package main

// LeetCode 3 - Longest Substring Without Repeating Characters
//
// Problem:
//   Given a string s, find the length of the longest substring without
//   repeating characters.
//
// Example:
//   Input:  s = "substring"
//   Output: 8
//
//   Explanation:
//   Track a count of each character inside the window [start, end].
//   When a character's count exceeds 1, the window has a duplicate, so
//   shrink from the left (decrementing counts) until it's unique again.
//   "substring" -> dropping the leading 's' leaves "ubstring" (8 unique chars).
//
// Pseudo code:
//   state[s[end]]++
//   if state[s[end]] > 1: state[s[start]]--; start++
//   maxLength = max(maxLength, end-start+1)

func main() {
	s := "substring"
	res := longestSubStr(s)
	println(res) // 8
}

func longestSubStr(s string) int {
	start := 0
	state := make(map[byte]int)
	maxLength := 0

	for end := 0; end < len(s); end++ {
		state[s[end]]++
		if state[s[end]] > 1 {
			state[s[start]]--
			start++
		}

		if end-start+1 > maxLength {
			maxLength = end - start + 1
		}
	}

	return maxLength
}
