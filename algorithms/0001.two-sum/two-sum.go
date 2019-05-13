package problem0001

// TwoSum function is the implementation of the first problem on leetcode
func TwoSum(nums []int, target int) []int {
	result := make(map[int]int, len(nums))
	for index, val := range nums {
		if j, ok := result[target-val]; ok {
			return []int{j, index}
		}

		result[val] = index
	}

	return nil
}

/* {2, 15, 11, 7} */
/*  0, 1 , 2,  3 */
