package main

import (
	"fmt"

	problem0704 "go_algorithms/algorithms/0704.binary-search"
)

/* problem 0001 -> var nums = []int{2, 7, 11, 15}
var target = 9*/
var binarySearch = []int{2, 5}

func main() {
	/*fmt.Println(problem0001.TwoSum(nums, target))
	fmt.Println(problem0003.LengthOfLongestSubstring("abcdaaa"))
	problem0344.ReverseString([]byte("hello"))
	fmt.Println(problem0844.BackspaceCompare("ab#c", "ad#c"))*/
	fmt.Println(problem0704.BinarySearch(binarySearch, 5)) // {2, 5} 5 index => 1
}
