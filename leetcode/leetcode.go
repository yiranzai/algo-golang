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
	return 0
}
