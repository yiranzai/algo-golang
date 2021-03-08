package offer

// offer1

//FindRepeatNumber https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/
func FindRepeatNumber(nums []int) int {
	bitmap := make(map[int]bool)
	for _, num := range nums {
		if bitmap[num] {
			return num
		}
		bitmap[num] = true
	}
	return -1
}

func FindRepeatNumber2(nums []int) int {
	for i := 0; i < len(nums); {
		if nums[i] == i {
			i++
			continue
		} else {
			if nums[nums[i]] == nums[i] {
				return nums[i]
			} else {
				nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
			}
		}
	}
	return -1
}


// https://leetcode-cn.com/problems/search-a-2d-matrix-ii
func findNumberIn2DArray(matrix [][]int, target int) bool {
	return false
}
