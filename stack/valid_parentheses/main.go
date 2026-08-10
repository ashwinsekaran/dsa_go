package main

// LeetCode 20 - Valid Parentheses
//
// Problem:
//   Given a string containing just the characters '(', ')', '{', '}',
//   '[' and ']', determine if the input string is valid — every open
//   bracket is closed by the same type of bracket, in the correct order.
//
// Example:
//   Input:  s = "({({}))}"
//   Output: false
//
//   Explanation:
//   Every closing bracket matches the most recently opened bracket
//   (LIFO order), so a stack tracks open brackets: push on open, pop
//   and compare on close. After "({({}" the stack is (,{,( and the
//   next ')' correctly pops '('; the following ')' then expects '('
//   but finds '{' on top — a mismatch — so the string is invalid.
//
// Pseudo code:
//   for each char in s:
//     if open bracket: push onto stack
//     if close bracket:
//       if stack empty or top doesn't match: return false
//       pop stack
//   return stack is empty

func main() {
	s := "({({}))}"
	println(isValid(s)) // false
}

func isValid(s string) bool {
	stack := []rune{}
	for _, char := range s {
		switch char {
		case '(', '{', '[':
			stack = append(stack, char)
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0

}
