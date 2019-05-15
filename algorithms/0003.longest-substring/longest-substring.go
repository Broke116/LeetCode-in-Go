package problem0003

// LengthOfLongestSubstring is the solution of problem 3
// dictionary initialization should be done inside the for loop. if it is done outside the function, then the time complexity increases. In addtion if 256 sized array is used, then the huge chunk of memory
// is allocated for nothing
func LengthOfLongestSubstring(s string) int {
	if s == " " {
		return 1
	}
	dict := make(map[string]int)
	max, left := 0, 0

	for i, v := range s {
		if _, ok := dict[string(v)]; !ok {
			dict[string(v)] = -1
		}
		if dict[string(v)] >= left {
			left = dict[string(v)] + 1
		} else if i+1-left > max {
			max = i + 1 - left
		}
		dict[string(v)] = i
	}
	return max
}

/*
"abcabcbb" 3
"bbbbb" 1
"pwwkew" 3
" " 1
"au" 2
"c" 1
"aab" 2

*/
