# dsa_go

DSA patterns and algorithms in Go — 8 problems, each in its own folder as a standalone `package main`.

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
