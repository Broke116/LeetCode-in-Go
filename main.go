package main

import (
	"fmt"

	problem0844 "go_algorithms/algorithms/0844.backspace-string-compare"
)

var nums = []int{2, 7, 11, 15}
var target = 9

func main() {
	//fmt.Println(problem0001.TwoSum(nums, target))
	// fmt.Println(problem0003.LengthOfLongestSubstring("abcdaaa"))
	// fmt.Println(problem0003.LengthOfLongestSubstring("bbbbbbbb"))
	// fmt.Println(problem0003.LengthOfLongestSubstring("aab"))
	// fmt.Println(problem0003.LengthOfLongestSubstring("au"))
	// problem0344.ReverseString([]byte("hello"))
	/*fmt.Println(problem0844.BackspaceCompare("ab#c", "ad#c"))
	fmt.Println(problem0844.BackspaceCompare("ab##", "c#d#"))
	fmt.Println(problem0844.BackspaceCompare("a##c", "#a#c"))
	fmt.Println(problem0844.BackspaceCompare("a#c", "b"))*/
	// fmt.Println(problem0844.BackspaceCompare("a##c", "#a#c"))
	fmt.Println(problem0844.BackspaceCompare("ab##", "c#d#"))
}
