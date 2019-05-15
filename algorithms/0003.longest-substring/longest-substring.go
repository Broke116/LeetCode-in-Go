package problem0003

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
