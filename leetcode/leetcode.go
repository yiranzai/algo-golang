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
	dp := []int{^(1 << 32)}
	j := 0
	for _, num := range nums {
		if dp[j] < num {
			dp = append(dp, num)
			j++
		} else {
			l, r := 0, j+1
			for l < r {
				m := l + (r-l)/2
				if dp[m] < num {
					l = m + 1
				} else {
					r = m
				}
			}
			dp[l] = num
		}
	}
	return j
}

// https://leetcode-cn.com/problems/word-break/
var (
	dict      map[string]bool
	dp        map[int]int
	maxLength int
)

func wordBreak(s string, wordDict []string) bool {
	dp = make(map[int]int)
	dict = make(map[string]bool)
	for _, s2 := range wordDict {
		maxLength = math.MaxInt(len(s2), maxLength)
		dict[s2] = true
	}

	return deepWordBreak(s, -1)
}

func deepWordBreak(s string, last int) bool {
	length := len(s)
	if len(s) == 0 {
		return true
	}
	start := last + 1
	for i := 0; i < length; i++ {
		if i > maxLength {
			break
		}
		if dict[s[0:i+1]] {
			if dp[start+i] == 1 {
				return true
			}
			if dp[start+i] == 2 {
				continue
			}
			if deepWordBreak(s[i+1:], start+i) {
				dp[start+i] = 1
				return true
			}
		}
	}
	dp[last] = 2
	return false
}

func wordBreak2(s string, wordDict []string) bool {
	// f[i] 表示前i个字符是否可以被切分
	// f[i] = f[j] && s[j+1~i] in wordDict
	// f[0] = true
	// return f[len]

	if len(s) == 0 {
		return true
	}
	f := make([]bool, len(s)+1)
	f[0] = true
	max, dict := maxLen(wordDict)
	for i := 1; i <= len(s); i++ {
		l := 0
		if i-max > 0 {
			l = i - max
		}
		for j := l; j < i; j++ {
			if f[j] && inDict(s[j:i], dict) {
				f[i] = true
				break
			}
		}
	}
	return f[len(s)]
}

func maxLen(wordDict []string) (int, map[string]bool) {
	dict := make(map[string]bool)
	max := 0
	for _, v := range wordDict {
		dict[v] = true
		if len(v) > max {
			max = len(v)
		}
	}
	return max, dict
}

func inDict(s string, dict map[string]bool) bool {
	_, ok := dict[s]
	return ok
}

// https://leetcode-cn.com/problems/longest-common-subsequence/
func longestCommonSubsequence(text1 string, text2 string) int {
	var t1Len, t2Len int
	t1Len = len(text1)
	t2Len = len(text2)
	dp := make([][]int, t1Len+1)
	for i := 0; i <= 1; i++ {
		dp[i] = make([]int, t2Len+1)
	}
	k, l := 0, 1
	for i := 1; i <= t1Len; i++ {
		for j := 1; j <= t2Len; j++ {
			if text1[i-1] == text2[j-1] {
				dp[k][j] = dp[l][j-1] + 1
			} else {
				dp[k][j] = math.MaxInt(dp[l][j], dp[k][j-1])
			}
		}
		l, k = k, l
	}

	return dp[l][t2Len]
}

// https://leetcode-cn.com/problems/edit-distance/
func minDistance(word1 string, word2 string) int {
	var t1Len, t2Len int
	t1Len = len(word1)
	t2Len = len(word2)
	if t1Len == 0 {
		return t2Len
	}
	if t2Len == 0 {
		return t1Len
	}
	dp := make([][]int, t1Len+1)
	for i := 0; i <= 1; i++ {
		dp[i] = make([]int, t2Len+1)
		for j := 1; j <= t2Len; j++ {
			dp[i][j] = j
		}
	}
	k, l := 0, 1
	for i := 1; i <= t1Len; i++ {
		dp[k][0] = i
		for j := 1; j <= t2Len; j++ {
			if word1[i-1] == word2[j-1] {
				dp[k][j] = dp[l][j-1]
			} else {
				dp[k][j] = math.MinInt(dp[l][j], dp[k][j-1], dp[l][j-1]) + 1
			}
		}
		l, k = k, l
	}

	return dp[l][t2Len]
}

// https://leetcode-cn.com/problems/coin-change/
func coinChange(coins []int, amount int) int {
	dps := make(map[int]int, amount)
	return deepCoinChange(coins, amount, dps)
}

func deepCoinChange(coins []int, amount int, dps map[int]int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}

	if v, ok := dps[amount]; ok {
		return v
	}
	var min *int
	for i := 0; i < len(coins); i++ {
		a := deepCoinChange(coins, amount-coins[i], dps)
		if a != -1 {
			if min == nil {
				min = &a
				*min += 1
			} else {
				if *min > a+1 {
					*min = a + 1
				}
			}
		}
	}
	if min == nil {
		dps[amount] = -1
		return -1
	}
	dps[amount] = *min
	return dps[amount]
}

// https://leetcode-cn.com/problems/minimum-path-sum/

func minPathSum2(grid [][]int) int {
	dps := make([][]int, len(grid))
	for i, i2 := range grid {
		dps[i] = make([]int, len(i2))
		for i3 := range dps[i] {
			dps[i][i3] = -1
		}
	}
	return deepMinPathSum(grid, dps, 0, 0)
}

func deepMinPathSum(grid, dps [][]int, x, y int) int {
	if dps[y][x] > -1 {
		return dps[y][x]
	}

	min := grid[y][x]

	if y == len(grid)-1 && x == len(grid[y])-1 {

	} else if y == len(grid)-1 {
		min += deepMinPathSum(grid, dps, x+1, y)
	} else if x == len(grid[y])-1 {
		min += deepMinPathSum(grid, dps, x, y+1)
	} else if y < len(grid)-1 && x < len(grid[y])-1 {
		min += math.MinInt(deepMinPathSum(grid, dps, x, y+1), deepMinPathSum(grid, dps, x+1, y))
	}

	dps[y][x] = min
	return min
}

func minPathSum(grid [][]int) int {
	lenY := len(grid) - 1
	lenX := len(grid[0]) - 1
	for y := 0; y <= lenY; y++ {
		for x := 1; x <= lenX; x++ {
			if y == 0 {
				grid[y][x] += grid[y][x-1]
				continue
			}
			grid[y][x] += math.MinInt(grid[y-1][x], grid[y][x-1])
		}
		if y == lenY {
			continue
		}
		grid[y+1][0] += grid[y][0]
	}

	return grid[lenY][lenX]
}

//https://leetcode-cn.com/problems/unique-paths/
func uniquePaths(m int, n int) int {
	if m == 1 && n == 1 {
		return 1
	}
	lenY := m - 1
	lenX := n - 1
	dp := make([]int, n)
	dp[0] = 1
	for y := 0; y <= lenY; y++ {
		for x := 1; x <= lenX; x++ {
			if y == 0 {
				dp[x] = 1
				continue
			}
			dp[x] += dp[x-1]
		}
	}

	return dp[lenX]
}

//https://leetcode-cn.com/problems/unique-paths-ii/
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	lenY := len(obstacleGrid) - 1
	lenX := len(obstacleGrid[0]) - 1
	dp := make([]int, lenX+1)
	dp[0] = 1
	boolY0 := false
	for y := 0; y <= lenY; y++ {
		boolY := 0
		if obstacleGrid[y][0] == 1 {
			boolY0 = true
		}
		if boolY0 {
			dp[0] = 0
			boolY = 1
		}
		for x := 1; x <= lenX; x++ {
			b := obstacleGrid[y][x] == 1
			if b {
				boolY++
				dp[x] = 0
				continue
			}
			if y == 0 {
				if boolY > 0 {
					dp[x] = 0
				} else {
					dp[x] = 1
				}
				continue
			}
			dp[x] += dp[x-1]
		}
		if boolY == lenY+1 {
			return 0
		}
	}

	return dp[lenX]
}

/**
https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	slow := head
	node := head
	fast := head.Next
	for {
		// 一直遍历快指针直到下一个不同的值出现
		for fast != nil && fast.Val == slow.Val {
			fast = fast.Next
		}

		// 如果快慢指针不相同 且慢指针是head，且快指针动过了，移动head
		if slow.Val == head.Val && slow.Next != fast {
			head = fast
			node = fast
		}
		// 快指针跑完了
		if fast == nil {
			if node != nil {
				node.Next = nil
			}
			return head
		}

		// 移动快慢指针
		slow = fast
		fast = fast.Next

		// 移动node指针
		if (fast == nil || (fast != nil && slow.Val != fast.Val)) && slow != node {
			node.Next = slow
			node = slow
		}

		continue
	}
}

/**
* https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
*/
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	slow := head
	fast := head.Next
	for {
		// 一直遍历快指针直到下一个不同的值出现
		for fast != nil && fast.Val == slow.Val {
			fast = fast.Next
			continue
		}
		slow.Next = fast
		slow = fast
		if slow == nil {
			return head
		}
		continue
	}
}

/**
* https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return math.MaxInt(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

/**
 * https://leetcode-cn.com/problems/balanced-binary-tree/
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	b, _ := deepIsBalanced(root)
	return b
}

func deepIsBalanced(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}
	b, left := deepIsBalanced(root.Left)
	if !b {
		return false, 0
	}
	b, right := deepIsBalanced(root.Right)
	if !b {
		return false, 0
	}
	return left-right <= 1 && left-right >= -1, math.MaxInt(left, right) + 1
}

/**
 * https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	return 0
}
