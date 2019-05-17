package problem0704

// BinarySearch is the implementation of the problem 0704
func BinarySearch(nums []int, target int) int {
	index := (len(nums) / 2) - 1
	end := len(nums) - 1
	first := 0

	if nums[index] == target {
		return index
	}

	for first <= end {
		if nums[index] < target {
			first = index + 1
		} else if nums[index] == target {
			return index
		} else {
			end = index - 1
		}
		index = (first + end) / 2
	}

	return -1
}

/* {-1, 0, 3, 5, 9, 12} find 5 or 2 */
/* {2, 5} 5 */
