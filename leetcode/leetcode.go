package leetcode

import (
	"strings"

	"github.com/yiranzai/go-utils/math"
)

// https://leetcode-cn.com/problems/valid-anagram/
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sum := make(map[byte]int)
	for index, b := range s {
		sum[byte(b)]++
		sum[t[index]]--
	}

	for _, v := range sum {
		if v != 0 {
			return false
		}
	}

	return true
}

// https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/
func replaceSpace(s string) string {
	var result []string
	empty := ' '
	for _, i := range s {
		if i == empty {
			result = append(result, "%20")
		} else {
			result = append(result, string(i))
		}
	}
	return strings.Join(result, "")
}

// https://leetcode-cn.com/problems/triangle/
// 空间复杂度 O(n) 2n
func minimumTotal(triangle [][]int) int {
	mark = make([]int, len(triangle))
	currentY = make([]int, len(triangle))
	return sumWay(triangle, 0, 0)
}

var mark []int
var currentY []int

func sumWay(array [][]int, x, y int) int {
	if y == len(array)-1 {
		return array[y][x]
	}
	if currentY[x] == y && y > 0 {
		return mark[x]
	}
	mark[x] = array[y][x] + math.MinInt(sumWay(array, x, y+1), sumWay(array, x+1, y+1))
	currentY[x] = y
	return mark[x]
}

// https://leetcode-cn.com/problems/longest-consecutive-sequence/
//
func longestConsecutive(nums []int) int {
	numSet := map[int]bool{}
	// 预存有值的数据
	for _, num := range nums {
		numSet[num] = true
	}
	count := 0
	for num := range numSet {
		if !numSet[num-1] {
			currentNum := num
			currentCount := 1
			for numSet[currentNum+1] {
				currentNum++
				currentCount++
			}
			if count < currentCount {
				count = currentCount
			}
		}
	}
	return count
}

//https://leetcode-cn.com/problems/climbing-stairs/
func climbStairs(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	a := 1
	b := 2
	c := 2
	for i := 3; i <= n; i++ {
		c = a + b
		a, b = b, c
	}

	return c
}

// https://leetcode-cn.com/problems/jump-game/

func canJump(nums []int) bool {
	length := len(nums) - 1
	if length == 0 {
		return true
	}
	dps := make([]bool, length)
dd:
	for i := length - 1; i >= 0; i-- {
		if length-i <= nums[i] {
			dps[i] = true
			continue
		}
		for j := nums[i]; j > 0; j-- {
			if dps[i+j] {
				dps[i] = true
				continue dd
			}
		}
		dps[i] = false
	}

	return dps[0]
}

// https://leetcode-cn.com/problems/jump-game-ii/
func jump(nums []int) int {
	length := len(nums) - 1
	if length == 0 {
		return 0
	}
	dps := make([]int, length+1)

	for i := 0; i <= length; i++ {
		for j := 1; j <= nums[i]; j++ {
			if dps[i] == 0 || dps[i+j] == 0 || dps[i+j] > dps[i] {
				dps[i+j] = dps[i] + 1
			}

			if i+j >= length {
				return dps[length]
			}
		}
	}

	return dps[length]
}

// https://leetcode-cn.com/problems/palindrome-partitioning-ii/
func minCut(s string) int {
	return 0
}

// https://leetcode-cn.com/problems/longest-increasing-subsequence/

func lengthOfLIS(nums []int) int {
	return 0
}

// https://leetcode-cn.com/problems/word-break/
func wordBreak(s string, wordDict []string) bool {
	return false
}
