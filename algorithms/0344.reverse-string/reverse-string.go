package problem0344

// ReverseString function is the implementation of the first problem on leetcode
func ReverseString(s []byte) {
	i := 0
	j := len(s) - 1
	for i < j {
		temp := s[j]
		s[j] = s[i]
		s[i] = temp
		j--
		i++
	}
}

/*
Input: ["h","e","l","l","o"]
Output: ["o","l","l","e","h"]

Input: ["H","a","n","n","a","h"]*
Output: ["h","a","n","n","a","H"]*/
