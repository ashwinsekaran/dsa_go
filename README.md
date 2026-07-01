# sliding_window_go

5 sliding window problems, each in its own folder as a standalone `package main`.

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
| — | Maximum Sum Subarray of Size K | Easy | [slidingwindow/maxium_sum_subarray/main.go](slidingwindow/maxium_sum_subarray/main.go) | Fixed window |
| 1423 | Maximum Points You Can Obtain from Cards | Medium | [slidingwindow/maximum_points_cards/main.go](slidingwindow/maximum_points_cards/main.go) | Fixed window on complement |
| 2461 | Maximum Sum of Distinct Subarrays With Length K | Medium | [slidingwindow/maximum_distinct_sum_subarray/main.go](slidingwindow/maximum_distinct_sum_subarray/main.go) | Fixed window + freq map |
| 3 | Longest Substring Without Repeating Characters | Medium | [dynamic_slidingwindow/longest_substr_without_repeat/main.go](dynamic_slidingwindow/longest_substr_without_repeat/main.go) | Variable window |
| 904 | Fruit Into Baskets | Medium | [dynamic_slidingwindow/variable_sliding_window_fruits_baskets/main.go](dynamic_slidingwindow/variable_sliding_window_fruits_baskets/main.go) | Variable window |

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
