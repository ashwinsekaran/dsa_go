# dsa_go

DSA patterns and algorithms in Go — 13 problems, each in its own folder as a standalone `package main`.

**Run any problem:**
```bash
go run <folder>/<subfolder>/main.go
# e.g.
go run slidingwindow/maxium_sum_subarray/main.go
```

---

## Problems

| # | Problem | Difficulty | File | Approach |
|---|---------|------------|------|----------|
| — | Max Subarray | Easy | [slidingwindow/maxium_sum_subarray/main.go](slidingwindow/maxium_sum_subarray/main.go) | Fixed window |
| 1423 | Card Points | Medium | [slidingwindow/maximum_points_cards/main.go](slidingwindow/maximum_points_cards/main.go) | Fixed window on complement |
| 2461 | Distinct Subarray | Medium | [slidingwindow/maximum_distinct_sum_subarray/main.go](slidingwindow/maximum_distinct_sum_subarray/main.go) | Fixed window + freq map |
| 3 | Longest Substring | Medium | [dynamic_slidingwindow/longest_substr_without_repeat/main.go](dynamic_slidingwindow/longest_substr_without_repeat/main.go) | Variable window |
| 904 | Fruit Baskets | Medium | [dynamic_slidingwindow/variable_sliding_window_fruits_baskets/main.go](dynamic_slidingwindow/variable_sliding_window_fruits_baskets/main.go) | Variable window |
| 11 | Container With Water | Medium | [2pointers/container_with_most_water/main.go](2pointers/container_with_most_water/main.go) | Two pointers |
| — | Pair Target Sum | Easy | [2pointers/pair_nums_to_find_a_sum/main.go](2pointers/pair_nums_to_find_a_sum/main.go) | Two pointers on sorted array |
| 252 | Meeting Rooms | Easy | [intervals/attend_meetings/main.go](intervals/attend_meetings/main.go) | Sort by start, check overlap |
| 20 | Valid Parentheses | Easy | [stack/valid_parentheses/main.go](stack/valid_parentheses/main.go) | Stack matching |
| 206 | Reverse Linked List | Easy | [linked_list/reverse_ll/main.go](linked_list/reverse_ll/main.go) | Iterative pointer reversal |
| 141 | Linked List Cycle | Easy | [linked_list/has_cycle/main.go](linked_list/has_cycle/main.go) | Visited-set traversal |
| 234 | Palindrome Linked List | Easy | [linked_list/isPalindrome/main.go](linked_list/isPalindrome/main.go) | Find mid, reverse half, compare |
| 21 | Merge Two Sorted Lists | Easy | [linked_list/merge2ll/main.go](linked_list/merge2ll/main.go) | Splice smaller node each step |

---

## Pseudo Code Summaries

### Maximum Sum Subarray of Size K
```
windowSum += nums[end]
if window size == k:
  maxSum = max(maxSum, windowSum)
  windowSum -= nums[start]; start++
```

### Maximum Points You Can Obtain from Cards
```
total = sum(cards)
windowSum = sum of first (len(cards)-k) cards   # the complement window
slide window of size len(cards)-k across cards
maxSum = max(maxSum, total - windowSum)         # picks = total - min complement
```

### Maximum Sum of Distinct Subarrays With Length K
```
windowSum += nums[end]; freqMap[nums[end]]++
if window size == k:
  if len(freqMap) == k:   # all elements distinct
    maxDistinctSum = max(maxDistinctSum, windowSum)
  windowSum -= nums[start]; decrement/delete freqMap[nums[start]]; start++
```

### Longest Substring Without Repeating Characters
```
state[s[end]]++
if state[s[end]] > 1:
  state[s[start]]--; start++
maxLength = max(maxLength, end-start+1)
```

### Fruit Into Baskets
```
basket[fruits[end]]++
if len(basket) > 2:
  basket[fruits[start]]--; delete if zero; start++
maxFruit = max(maxFruit, end-start+1)
```

### Container With Most Water
```
left, right = 0, len(heights)-1
area = min(heights[left], heights[right]) * (right - left)
advance the pointer at the shorter height
return max area seen
```

### Pair with Target Sum
```
sort(nums)
left, right = 0, len(nums)-1
while left < right:
  sum = nums[left] + nums[right]
  if sum == target: return true
  if sum > target: right--
  else: left++
return false
```

### Meeting Rooms
```
sort intervals by start time
for i in 1..n-1:
  if intervals[i].start < intervals[i-1].end: return false
return true
```

### Valid Parentheses
```
for each char in s:
  if open bracket: push onto stack
  if close bracket:
    if stack empty or top doesn't match: return false
    pop stack
return stack is empty
```

### Reverse Linked List
```
prev = nil; current = head
while current != nil:
  next = current.Next
  current.Next = prev
  prev = current
  current = next
return prev
```

### Linked List Cycle
```
visited = {}
for node := head; node != nil; node = node.Next:
  if visited[node]: return true
  visited[node] = true
return false
```

### Palindrome Linked List
```
slow, fast = head, head
while fast != nil && fast.Next != nil:
  slow = slow.Next; fast = fast.Next.Next
reverse the list starting at slow -> prev
walk head and prev together, comparing values; mismatch -> false
```

### Merge Two Sorted Lists
```
pick smaller head as start; advance that list
while both lists non-empty:
  attach smaller of l1.Val, l2.Val; advance that list
attach whichever list remains
return start
```
