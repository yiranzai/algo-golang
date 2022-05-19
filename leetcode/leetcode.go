package leetcode

import (
	"container/heap"
	"container/list"
	"sort"
	"strings"

	"github.com/yiranzai/go-utils/leetcode"
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
	dp := []int{INVALID}
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
 *     Next *leetcode.ListNode
 * }
*/

func deleteDuplicates2(head *leetcode.ListNode) *leetcode.ListNode {
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
 *     Next *leetcode.ListNode
 * }
*/
func deleteDuplicates(head *leetcode.ListNode) *leetcode.ListNode {
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
 *     Left *leetcode.TreeNode
 *     Right *leetcode.TreeNode
 * }
*/

func maxDepth(root *leetcode.TreeNode) int {
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
 *     Left *leetcode.TreeNode
 *     Right *leetcode.TreeNode
 * }
 */
func isBalanced(root *leetcode.TreeNode) bool {
	b, _ := deepIsBalanced(root)
	return b
}

func deepIsBalanced(root *leetcode.TreeNode) (bool, int) {
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
 *     Left *leetcode.TreeNode
 *     Right *leetcode.TreeNode
 * }
 */
func maxPathSum(root *leetcode.TreeNode) int {
	var sum int
	sum = INVALID
	deepMaxPathSum(root, &sum)
	return sum
}

const INVALID = ^(1 << 32)

func deepMaxPathSum(root *leetcode.TreeNode, sum *int) int {
	if root == nil {
		return INVALID
	}
	rootVal := root.Val
	left := deepMaxPathSum(root.Left, sum)
	right := deepMaxPathSum(root.Right, sum)

	// 找出左右及单独的自己中最大的
	a := math.MaxInt(left+rootVal, rootVal, right+rootVal)
	*sum = math.MaxInt(*sum, a, left+rootVal+right)
	return a
}

/**
* https://leetcode-cn.com/problems/rotate-list/
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *leetcode.ListNode
 * }
*/
func rotateRight(head *leetcode.ListNode, k int) *leetcode.ListNode {
	if head == nil || k == 0 {
		return head
	}
	node := head
	length := 1
	for node.Next != nil {
		node = node.Next
		length++
	}
	k = k % length
	if k == 0 {

		return head
	}
	last := node
	node = head
	for i := 0; i < length-k-1; i++ {
		node = node.Next
	}
	last.Next = head
	head = node.Next
	node.Next = nil
	return head
}

/**

https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
*/
func lowestCommonAncestor(root, p, q *leetcode.TreeNode) *leetcode.TreeNode {
	node, _, _ := deepLowestCommonAncestor(root, p, q)
	return node
}

func deepLowestCommonAncestor(root, p, q *leetcode.TreeNode) (*leetcode.TreeNode, bool, bool) {
	if root == nil {
		return nil, false, false
	}
	// 左边俩都搜到了
	left, blp, blq := deepLowestCommonAncestor(root.Left, p, q)
	if blp && blq {
		return left, blp, blq
	}

	// 右边俩都搜到了
	right, brp, brq := deepLowestCommonAncestor(root.Right, p, q)
	if brp && brq {
		return right, brp, brq
	}

	rp, rq := root.Val == p.Val, root.Val == q.Val

	// 左右及根都没有
	if !rp && !rq && !brp && !brq && !blp && !blq {
		return nil, false, false
	}

	// 左边或者右边搜索到其中一个
	if (rp && (blq || brq)) || (rq && (blp || brp)) || (blq && brp) || (blp && brq) {
		return root, true, true
	}

	return root, rp || blp || brp, rq || blq || brq
}

/**
 * https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *leetcode.TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var list []*leetcode.TreeNode
	list = append(list, root)
	res = make([][]int, 0, 0)
	length := len(list)
	for length > 0 {
		b := make([]int, length)
		for i := 0; i < length; i++ {
			b[i] = list[i].Val
			if list[i].Left != nil {
				list = append(list, list[i].Left)
			}
			if list[i].Right != nil {
				list = append(list, list[i].Right)
			}
		}
		list = list[length:]
		length = len(list)
		res = append(res, b)
	}
	return res
}

/**
 * https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *leetcode.TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var list []*leetcode.TreeNode
	list = append(list, root)
	res = make([][]int, 0, 0)
	length := len(list)
	for length > 0 {
		b := make([]int, length)
		for i := 0; i < length; i++ {
			b[i] = list[i].Val
			if list[i].Left != nil {
				list = append(list, list[i].Left)
			}
			if list[i].Right != nil {
				list = append(list, list[i].Right)
			}
		}
		list = list[length:]
		length = len(list)
		res = append([][]int{b}, res...)
	}
	return res
}

/**
 * https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *leetcode.TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var list []*leetcode.TreeNode
	list = append(list, root)
	res = make([][]int, 0, 0)
	length := len(list)
	mark := true
	for length > 0 {
		b := make([]int, length)
		for i := 0; i < length; i++ {
			if mark {
				b[i] = list[i].Val
			} else {
				b[i] = list[length-1-i].Val
			}
			if list[i].Left != nil {
				list = append(list, list[i].Left)
			}
			if list[i].Right != nil {
				list = append(list, list[i].Right)
			}
		}

		mark = !mark
		list = list[length:]
		length = len(list)
		res = append(res, b)
	}
	return res
}

/**
 * https://leetcode-cn.com/problems/validate-binary-search-tree/submissions/
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *leetcode.TreeNode) bool {
	bb, _, _ := deepIsValidBST(root)
	return bb == 1
}

func deepIsValidBST(root *leetcode.TreeNode) (a int, max int, min int) {
	if root == nil {
		return 2, 0, 0
	}

	if root.Left == nil && root.Right == nil {
		return 1, root.Val, root.Val
	}

	bl, lMax, lMin := deepIsValidBST(root.Left)
	if bl == 0 {
		return bl, 0, 0
	}

	br, rMax, rMin := deepIsValidBST(root.Right)
	if br == 0 {
		return br, 0, 0
	}

	if bl != 2 && root.Val <= lMax {
		return 0, 0, 0
	}

	if br != 2 && root.Val >= rMin {
		return 0, 0, 0
	}

	if bl == 2 {
		return 1, math.MaxInt(rMax, root.Val), math.MinInt(rMin, root.Val)
	}

	if br == 2 {
		return 1, math.MaxInt(lMax, root.Val), math.MinInt(lMin, root.Val)
	}

	return 1, math.MaxInt(lMax, rMax, root.Val), math.MinInt(root.Val, lMin, rMin)
}

/**
 * https://leetcode-cn.com/problems/binary-search-tree-iterator/
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	root *leetcode.TreeNode
	list []*leetcode.TreeNode
}

func BSTConstructor(root *leetcode.TreeNode) BSTIterator {
	res := BSTIterator{root: root}
	res.DeepLeft(root)
	return res
}

func (this *BSTIterator) DeepLeft(root *leetcode.TreeNode) {
	for root != nil {
		this.list = append(this.list, root)
		root = root.Left
	}
}

func (this *BSTIterator) Next() int {
	t := this.list[len(this.list)-1]
	this.list = this.list[0 : len(this.list)-1]
	if t.Right != nil || !this.HasNext() {
		this.DeepLeft(t.Right)
	}
	return t.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.list) > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := BSTConstructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

// https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoBST(root *leetcode.TreeNode, val int) *leetcode.TreeNode {
	if root == nil {
		return &leetcode.TreeNode{Val: val}
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	}
	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}

/**
 * https://leetcode-cn.com/problems/reverse-linked-list/
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *leetcode.ListNode) *leetcode.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	node := head.Next
	head.Next = nil
	for node != nil {
		head, node.Next = node.Next, head
		head, node = node, head
	}
	return head
}

func deepReverseList(head *leetcode.ListNode) *leetcode.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	res := deepReverseList(head.Next)
	head.Next = nil
	node := res
	for node.Next != nil {
		node = node.Next
	}
	node.Next = head
	return res
}

/**
* https://leetcode-cn.com/problems/reverse-linked-list-ii/
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
*/
func reverseBetween(head *leetcode.ListNode, left int, right int) *leetcode.ListNode {
	if head == nil || head.Next == nil || right < 2 || left == right {
		return head
	}
	var end *leetcode.ListNode
	var tEnd *leetcode.ListNode
	var tNode *leetcode.ListNode
	var tHead *leetcode.ListNode
	current := head
	for i := 1; i <= right-1; i++ {
		// 记录left前的节点
		if i == left-1 {
			end = current
		}
		if i < left-1 {
			current = current.Next
		}
		// 在left时做初始化操作
		if i == left {
			if end == nil {
				tHead = head
			} else {
				tHead = end.Next
			}
			tNode = tHead.Next
			tHead.Next = nil
			tEnd = tHead
		}
		// 开始搞事
		if i >= left {
			tHead, tNode.Next = tNode.Next, tHead
			tHead, tNode = tNode, tHead
		}
	}
	tEnd.Next = tNode
	if end == nil {
		return tHead
	}
	end.Next = tHead
	return head
}

// https://leetcode-cn.com/problems/merge-two-sorted-lists/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *leetcode.ListNode, l2 *leetcode.ListNode) *leetcode.ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var head *leetcode.ListNode
	var node *leetcode.ListNode
	if l1.Val <= l2.Val {
		head = l1
		l1 = l1.Next
	} else {
		head = l2
		l2 = l2.Next
	}
	node = head
	for l1 != nil && l2 != nil {
		for l1 != nil && l1.Val <= l2.Val {
			node.Next = l1
			node = node.Next
			l1 = l1.Next
		}
		for l1 != nil && l2 != nil && l1.Val > l2.Val {
			node.Next = l2
			node = node.Next
			l2 = l2.Next
		}
	}

	if l1 == nil {
		node.Next = l2
	}
	if l2 == nil {
		node.Next = l1
	}

	return head
}

// https://leetcode-cn.com/problems/partition-list/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *leetcode.ListNode, x int) *leetcode.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	left := head
	right := head.Next
	var newHead *leetcode.ListNode

	if head.Val < x {
		newHead, head = head, head.Next
	}
	bNode := newHead
	for right != nil {
		if right.Val < x {
			if bNode == nil {
				newHead = right
				bNode = newHead
			} else {
				bNode.Next = right
				bNode = bNode.Next
			}
			// 当前是head,head也移动
			if right == head {
				head = right.Next
			} else {
				left.Next = right.Next
			}
		}
		left, right = left.Next, right.Next
	}
	if bNode == nil {
		return head
	}
	bNode.Next = head
	return newHead
}

func partition2(head *leetcode.ListNode, x int) *leetcode.ListNode {
	small := &leetcode.ListNode{}
	smallHead := small
	large := &leetcode.ListNode{}
	largeHead := large
	for head != nil {
		if head.Val < x {
			small.Next = head
			small = small.Next
		} else {
			large.Next = head
			large = large.Next
		}
		head = head.Next
	}
	large.Next = nil
	small.Next = largeHead.Next
	return smallHead.Next
}

/**
 * https://leetcode-cn.com/problems/sort-list/
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *leetcode.ListNode) *leetcode.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var end *leetcode.ListNode
	var l *leetcode.ListNode
	var m *leetcode.ListNode
	var r *leetcode.ListNode
	var last *leetcode.ListNode
	for head != end {
		l = nil
		m, r = head, head.Next
		max := m.Val
		mark := true
		for r != end {
			if max < r.Val {
				max = r.Val
			} else if max > r.Val {
				mark = false
				m.Next, r.Next = r.Next, m
				if l == nil {
					head = r
				} else {
					l.Next = r
				}
				r, m = m, r
				last = m.Next
			}
			if l == nil {
				l = head
			} else {
				l = l.Next
			}
			m, r = m.Next, r.Next
		}
		if mark == true {
			break
		}
		end = last
	}
	return head
}

//https://leetcode-cn.com/problems/reverse-bits/
func reverseBits(num uint32) uint32 {
	var sum uint32
	for i := 0; i < 32 && num > 0; i++ {
		sum += (num & 1) << (31 - i)
		num >>= 1
	}
	return sum
}

const (
	m1 = 0x55555555 // 01010101010101010101010101010101
	m2 = 0x33333333 // 00110011001100110011001100110011
	m4 = 0x0f0f0f0f // 00001111000011110000111100001111
	m8 = 0x00ff00ff // 00000000111111110000000011111111
)

func reverseBits2(n uint32) uint32 {
	n = n>>1&m1 | n&m1<<1
	n = n>>2&m2 | n&m2<<2
	n = n>>4&m4 | n&m4<<4
	n = n>>8&m8 | n&m8<<8
	return n>>16 | n<<16
}

/**
 * https://leetcode-cn.com/problems/reorder-list/
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *leetcode.ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}
	list := make([]*leetcode.ListNode, 0)
	node := head
	for node != nil {
		list = append(list, node)
		node = node.Next
	}
	var i int
	for i = 1; i <= len(list)>>1; i++ {
		end := list[len(list)-i:][0]
		list[i-1].Next, end.Next = end, list[i]
	}
	list[i-1].Next = nil
}

func reverseList2(head *leetcode.ListNode) *leetcode.ListNode {
	// 1->2->3
	if head == nil {
		return head
	}
	var prev *leetcode.ListNode
	for head != nil {
		t := head.Next
		head.Next = prev
		prev = head
		head = t
	}
	return prev
}

/**
 * https://leetcode-cn.com/problems/linked-list-cycle/
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *leetcode.ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	if head.Next.Next == head {
		return true
	}
	fast := head.Next.Next
	slow := head

	for slow != nil && fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow, fast = slow.Next, fast.Next.Next
	}
	return false
}

/**
 * https://leetcode-cn.com/problems/linked-list-cycle-ii/
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *leetcode.ListNode) *leetcode.ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	if head.Next.Next == head {
		return head
	}
	fast := head.Next
	slow := head
	first := true
	for slow != nil && fast != nil && fast.Next != nil {
		if slow == fast {
			if first {
				first = false
				slow = head
				fast = fast.Next
				if fast == slow {
					return slow
				}
			} else {
				return slow
			}
		}
		if first {
			slow, fast = slow.Next, fast.Next.Next
		} else {
			slow, fast = slow.Next, fast.Next
		}
	}
	return nil
}

func detectCycle2(head *leetcode.ListNode) *leetcode.ListNode {
	// 思路：快慢指针，快慢相遇之后，慢指针回到头，快慢指针步调一致一起移动，相遇点即为入环点
	if head == nil {
		return head
	}
	fast := head.Next
	slow := head
	first := true
	for fast != nil && fast.Next != nil {
		if fast == slow {
			// 慢指针重新从头开始移动，快指针从第一次相交点下一个节点开始移动
			if first {
				slow = head
				fast = fast.Next // 注意
				if fast == slow {
					return slow
				}
				first = false
			} else {
				return slow
			}
		}
		if first {
			fast = fast.Next.Next
			slow = slow.Next
		} else {
			fast = fast.Next
			slow = slow.Next
		}
	}
	return nil
}

// https://leetcode-cn.com/problems/search-a-2d-matrix/
func searchMatrix(matrix [][]int, target int) bool {
	length := len(matrix)
	left := 0
	right := length
	for left < right {
		// 避免数据超范围
		mid := left + ((right - left) >> 1)
		if matrix[mid][0] == target {
			return true
		} else if matrix[mid][0] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	index := left - 1
	if index < 0 {
		return false
	}
	length = len(matrix[index])
	left = 0
	right = length
	for left < right {
		// 避免数据超范围
		mid := left + ((right - left) >> 1)
		if matrix[index][mid] == target {
			return true
		} else if matrix[index][mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return false
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *leetcode.ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	tail := reverseList2(slow.Next)
	slow.Next = nil
	for head != nil && tail != nil {
		if head.Val != tail.Val {
			return false
		}
		head, tail = head.Next, tail.Next
	}

	return true
}

func isPalindrome2(head *leetcode.ListNode) bool {
	var recursivelyCheck func(*leetcode.ListNode) bool
	recursivelyCheck = func(curNode *leetcode.ListNode) bool {
		if curNode != nil {
			if !recursivelyCheck(curNode.Next) {
				return false
			}
			if curNode.Val != head.Val {
				return false
			}
			head = head.Next
		}
		return true
	}
	return recursivelyCheck(head)
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

type RandomNode struct {
	Val    int
	Next   *RandomNode
	Random *RandomNode
}

func copyRandomList(head *RandomNode) *RandomNode {
	if head == nil {
		return head
	}
	node := head
	for node != nil {
		node.Next, node = &RandomNode{Next: node.Next, Val: node.Val, Random: node.Random}, node.Next
	}
	node = head.Next
	for node != nil {
		if node.Random != nil {
			node.Random = node.Random.Next
		}
		if node.Next == nil {
			break
		}
		node = node.Next.Next
	}
	node = head.Next
	b := head.Next
	for head != nil && node.Next != nil {
		head, head.Next = head.Next.Next, head.Next.Next
		node, node.Next = node.Next.Next, node.Next.Next
	}
	head.Next = nil
	return b
}

// https://leetcode-cn.com/problems/single-number/
func singleNumber1(nums []int) int {
	for i := 1; i < len(nums); i++ {
		nums[0] ^= nums[i]
	}
	return nums[0]
}

// https://leetcode-cn.com/problems/single-number-ii/
func singleNumber(nums []int) int {
	a, b := 0, 0
	for _, num := range nums {
		b = (^a) & (b ^ num)
		a = (^b) & (a ^ num)
	}
	return b
}

// https://leetcode-cn.com/problems/single-number-iii/
func singleNumber3(nums []int) []int {
	diff := 0
	for i := 0; i < len(nums); i++ {
		diff ^= nums[i]
	}
	result := []int{diff, diff}
	// 去掉末尾的1后异或diff就得到最后一个1的位置
	diff = (diff & (diff - 1)) ^ diff
	for i := 0; i < len(nums); i++ {
		if diff&nums[i] == 0 {
			result[0] ^= nums[i]
		} else {
			result[1] ^= nums[i]
		}
	}
	return result
}

// https://leetcode-cn.com/problems/subsets-ii/submissions/
var res [][]int

func subsetsWithDup(nums []int) [][]int {
	res = make([][]int, 0)
	sort.Ints(nums)
	dfs([]int{}, nums, 0)
	return res
}

func dfs(temp, nums []int, start int) {
	tmp := make([]int, len(temp))
	copy(tmp, temp)
	res = append(res, tmp)
	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] { // skip
			continue
		}
		temp = append(temp, nums[i])
		dfs(temp, nums, i+1)
		temp = temp[:len(temp)-1]
	}
}

//https://leetcode-cn.com/problems/number-of-1-bits/
func hammingWeight(num uint32) int {
	sum := 0
	for i := 0; i < 32; i++ {
		if num&1 == 1 {
			sum++
		}
		num = num >> 1
	}
	return sum
}

// https://leetcode-cn.com/problems/counting-bits/
func countBits(num int) []int {
	res := make([]int, 0)
	thisOne := 1
	nextOne := 2
	for i := 0; i <= num; i++ {
		if i == 0 {
			res = append(res, 0)
		} else if i == nextOne {
			res = append(res, 1)
			nextOne = nextOne << 1
			thisOne = thisOne << 1
		} else {
			res = append(res, res[i-thisOne]+1)
		}
	}
	return res
}

// https://leetcode-cn.com/problems/bitwise-and-of-numbers-range/
func rangeBitwiseAnd(left int, right int) int {
	if left == right {
		return left & right
	}

	if right >= left<<2 {
		return 0
	}
	end := 0

	for i := 0; i < 32; i++ {
		if left <= end && right > end {
			return 0
		}
		if end >= right {
			e := end >> 1
			t := end ^ e
			return t | (rangeBitwiseAnd(left&e, right&e))
		}
		end = (end << 1) + 1
	}
	return 0
}

func rangeBitwiseAnd2(m int, n int) int {
	for m < n {
		n &= (n - 1)
	}
	return n
}

// LRUCache1 https://leetcode-cn.com/problems/lru-cache/
type LRUCache1 struct {
	limit  int
	list   *list.List
	eMap   map[int]*list.Element
	values map[int]int
}

func LRUConstructor(capacity int) LRUCache1 {
	return LRUCache1{limit: capacity, list: list.New(), eMap: make(map[int]*list.Element), values: make(map[int]int)}
}

func (this *LRUCache1) Get(key int) int {
	if e, ok := this.eMap[key]; ok {
		this.list.Remove(e)
		front := this.list.PushFront(key)
		this.eMap[key] = front
		return this.values[key]
	}
	return -1
}

func (this *LRUCache1) Put(key int, value int) {
	length := this.list.Len()
	if e, ok := this.eMap[key]; ok {
		this.list.Remove(e)
	} else if length >= this.limit {
		back := this.list.Back()
		this.list.Remove(back)
		delete(this.eMap, back.Value.(int))
		delete(this.values, back.Value.(int))
	}
	front := this.list.PushFront(key)
	this.eMap[key] = front
	this.values[key] = value
}

// 分割线

type LRUCache struct {
	head, tail *Node
	keys       map[int]*Node
	capacity   int
}

type Node struct {
	key, val   int
	prev, next *Node
}

func ConstructorLRU(capacity int) LRUCache {
	return LRUCache{keys: make(map[int]*Node), capacity: capacity}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.keys[key]; ok {
		this.Remove(node)
		this.Add(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.keys[key]; ok {
		node.val = value
		this.Remove(node)
		this.Add(node)
		return
	} else {
		node = &Node{key: key, val: value}
		this.keys[key] = node
		this.Add(node)
	}
	if len(this.keys) > this.capacity {
		delete(this.keys, this.tail.key)
		this.Remove(this.tail)
	}
}

func (this *LRUCache) Add(node *Node) {
	node.prev = nil
	node.next = this.head
	if this.head != nil {
		this.head.prev = node
	}
	this.head = node
	if this.tail == nil {
		this.tail = node
		this.tail.next = nil
	}
}

func (this *LRUCache) Remove(node *Node) {
	if node == this.head {
		this.head = node.next
		if node.next != nil {
			node.next.prev = nil
		}
		node.next = nil
		return
	}
	if node == this.tail {
		this.tail = node.prev
		node.prev.next = nil
		node.prev = nil
		return
	}
	node.prev.next = node.next
	node.next.prev = node.prev
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

// LFUCache https://leetcode-cn.com/problems/lfu-cache/
type LFUCache struct {
	keys     map[int]*LFUNode
	list     map[int]*LFUList
	capacity int
	min      int
}

type LFUList struct {
	head, tail *LFUNode
}

type LFUNode struct {
	key, val   int
	num        int
	prev, next *LFUNode
}

func ConstructorLFU(capacity int) LFUCache {
	return LFUCache{keys: make(map[int]*LFUNode), list: make(map[int]*LFUList), capacity: capacity}
}

func (this *LFUCache) Get(key int) int {
	if node, ok := this.keys[key]; ok {
		this.ListAdd(node)
		return node.val
	}
	return -1
}

func (this *LFUCache) ListAdd(node *LFUNode) {
	if l, ok := this.list[node.num]; ok {
		l.Remove(node)
		if node.num <= this.min && l.tail != nil {
			this.min = l.tail.num
		} else {
			for i := 0; i >= 0; i++ {
				if len(this.list) == 0 {
					this.min = 0
					break
				}
				if l, ok = this.list[i]; ok {
					if l.head != nil {
						this.min = l.head.num
						break
					} else {
						delete(this.list, i)
					}
				}
			}
		}
	}

	node.num++
	if l, ok := this.list[node.num]; ok {
		l.Add(node)
		if node.num <= this.min {
			this.min = l.tail.num
		}
	} else {
		this.list[node.num] = &LFUList{head: node, tail: node}
		if node.num <= this.min {
			this.min = node.num
		}
	}
}

func (this *LFUCache) ListRemove() {
	if l, ok := this.list[this.min]; ok {
		delete(this.keys, l.tail.key)
		l.Remove(l.tail)
		if l.head == nil {
			delete(this.list, this.min)
			this.min++
		}
	}
}

func (this *LFUCache) Put(key int, value int) {
	if node, ok := this.keys[key]; ok {
		node.val = value
		this.ListAdd(node)
	} else {
		if len(this.keys) >= this.capacity {
			this.ListRemove()
		}
		node = &LFUNode{key: key, val: value, num: 1}
		this.keys[key] = node
		if l, ok := this.list[node.num]; ok {
			l.Add(node)
			this.min = l.tail.num
		} else {
			this.list[node.num] = &LFUList{head: node, tail: node}
		}
	}
}

func (this *LFUList) Add(node *LFUNode) {
	node.prev = nil
	node.next = this.head
	if this.head != nil {
		this.head.prev = node
	}
	this.head = node
	if this.tail == nil {
		this.tail = node
		this.tail.next = nil
	}
}

func (this *LFUList) Remove(node *LFUNode) {
	if node == this.head && this.tail == node {
		this.head = nil
		this.tail = nil
		return
	}
	if node == this.head {
		this.head = node.next
		if node.next != nil {
			node.next.prev = nil
		}
		node.next = nil
		return
	}
	if node == this.tail {
		this.tail = node.prev
		node.prev.next = nil
		node.prev = nil
		return
	}
	node.prev.next = node.next
	node.next.prev = node.prev
}

// -----------------------

type LFUCache2 struct {
	nodes    map[int]*list.Element
	lists    map[int]*list.List
	capacity int
	min      int
}

type node struct {
	key       int
	value     int
	frequency int
}

// Constructor define
func LFUConstructor(capacity int) LFUCache2 {
	return LFUCache2{nodes: make(map[int]*list.Element),
		lists:    make(map[int]*list.List),
		capacity: capacity,
		min:      0,
	}
}

// Get define
func (lfuCache *LFUCache2) Get(key int) int {
	value, ok := lfuCache.nodes[key]
	if !ok {
		return -1
	}
	currentNode := value.Value.(*node)
	lfuCache.lists[currentNode.frequency].Remove(value)
	currentNode.frequency++
	if _, ok := lfuCache.lists[currentNode.frequency]; !ok {
		lfuCache.lists[currentNode.frequency] = list.New()
	}
	newList := lfuCache.lists[currentNode.frequency]
	newNode := newList.PushFront(currentNode)
	lfuCache.nodes[key] = newNode
	if currentNode.frequency-1 == lfuCache.min && lfuCache.lists[currentNode.frequency-1].Len() == 0 {
		lfuCache.min++
	}
	return currentNode.value
}

// Put define
func (lfuCache *LFUCache2) Put(key int, value int) {
	if lfuCache.capacity == 0 {
		return
	}
	if currentValue, ok := lfuCache.nodes[key]; ok {
		currentNode := currentValue.Value.(*node)
		currentNode.value = value
		lfuCache.Get(key)
		return
	}
	if lfuCache.capacity == len(lfuCache.nodes) {
		currentList := lfuCache.lists[lfuCache.min]
		backNode := currentList.Back()
		delete(lfuCache.nodes, backNode.Value.(*node).key)
		currentList.Remove(backNode)
	}
	lfuCache.min = 1
	currentNode := &node{
		key:       key,
		value:     value,
		frequency: 1,
	}
	if _, ok := lfuCache.lists[1]; !ok {
		lfuCache.lists[1] = list.New()
	}
	newList := lfuCache.lists[1]
	newNode := newList.PushFront(currentNode)
	lfuCache.nodes[key] = newNode
}

// -----------------------

type LFUCache3 struct {
	capacity int
	pq       PriorityQueue
	hash     map[int]*Item
	counter  int
}

func LFConstructor(capacity int) LFUCache3 {
	lfu := LFUCache3{
		pq:       PriorityQueue{},
		hash:     make(map[int]*Item, capacity),
		capacity: capacity,
	}
	return lfu
}

func (this *LFUCache3) Get(key int) int {
	if this.capacity == 0 {
		return -1
	}
	if item, ok := this.hash[key]; ok {
		this.counter++
		this.pq.update(item, item.value, item.frequency+1, this.counter)
		return item.value
	}
	return -1
}

func (this *LFUCache3) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}
	// fmt.Printf("Put %d\n", key)
	this.counter++
	// 如果存在，增加 frequency，再调整堆
	if item, ok := this.hash[key]; ok {
		this.pq.update(item, value, item.frequency+1, this.counter)
		return
	}
	// 如果不存在且缓存满了，需要删除。在 hashmap 和 pq 中删除。
	if len(this.pq) == this.capacity {
		item := heap.Pop(&this.pq).(*Item)
		delete(this.hash, item.key)
	}
	// 新建结点，在 hashmap 和 pq 中添加。
	item := &Item{
		value: value,
		key:   key,
		count: this.counter,
	}
	heap.Push(&this.pq, item)
	this.hash[key] = item
}

// An Item is something we manage in a priority queue.
type Item struct {
	value     int // The value of the item; arbitrary.
	key       int
	frequency int // The priority of the item in the queue.
	count     int // use for evicting the oldest element
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	if pq[i].frequency == pq[j].frequency {
		return pq[i].count < pq[j].count
	}
	return pq[i].frequency < pq[j].frequency
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value int, frequency int, count int) {
	item.value = value
	item.count = count
	item.frequency = frequency
	heap.Fix(pq, item.index)
}

// AcNode AC自动机的节点
type AcNode struct {
	next map[byte]int
	deep int
	over int
	fail int
}

// AC AC算法，全名Alfred-Corasick自动机算法，这是一种多模式匹配的算法。这种算法某种程度上可说是
// KMP算法的多模式版本。事实上，如果只给AC算法提供一个字符串来生成AC自动机，就会发现，AC自动机
// 的goto链就是模式串，而fail链正是next数组的AC自动机版本。作为KMP算法的扩展版本，AC算法主要
// 用于多模式匹配。AC算法的预处理过程包括两个部分，其一是根据多个模式串构造Trie树的构建，其二是
// 在Trie树上添加fail“指针”，亦即匹配失败时改变当前活动节点为其指向的节点。类型AC表示由一系列
// 节点构成的AC自动机。
type AC struct {
	State []AcNode
	Final []string
	Count int
}

// NewAC 创建一个AC自动机，使用模式串本身作为其命名
func NewAC(ts []string) *AC {
	this := new(AC)
	for _, t := range ts {
		this.Add(t)
	}
	this.Prepare()
	return this
}

// 清空结构体，以重新生成多模式匹配
func (this *AC) Reset() {
	this.State = nil
	this.Final = nil
	this.Count = 0
}

// 向AC自动机中添加模式串
func (this *AC) Add(temp string) {
	p := len(this.Final)
	this.Final = append(this.Final, temp)
	if this.Count == 0 {
		this.newAcNode()
	}
	this.incStr(p, 0, temp)
}

// 重新生成fail指针，每次添加完一到多个新的模式串之后，在匹配文本串之前，都应执行Prepare方法
func (this *AC) Prepare() {
	type cell struct {
		pre int
		chr byte
		now int
	}
	this.State[0].deep = 0
	stack := make([]cell, this.Count)
	for i, j := 0, 1; j < this.Count; i++ {
		t := stack[i].now
		p := &this.State[t]
		if p.next != nil {
			for k, v := range p.next {
				this.State[v].deep = p.deep + 1
				stack[j] = cell{t, k, v}
				j++
			}
		}
	}
	for i := 1; i < this.Count; i++ {
		v := stack[i]
		s := &this.State[v.now] // 当前节点
		if v.pre != 0 {
			p := &this.State[v.pre] // 当前节点的父节点
			t := p.fail             // 父节点的fail指针
			for {
				if t < 0 {
					s.fail = 0
					break
				}
				p = &this.State[t]
				t = p.fail
				if p.next != nil {
					if n, ok := p.next[v.chr]; ok {
						s.fail = n
						break
					}
				}
			}
		} else {
			s.fail = 0
		}
	}
	this.State[0].fail = 0
}

// 搜索第一个匹配的位置，返回起始位置下标和模式串的下标，没有返回-1，-1
func (this *AC) Index(s []byte) (int, int) {
	t := 0
	for i, l := 0, len(s); i < l; {
		p := &this.State[t]
		if p.over >= 0 {
			return i - p.deep, p.over
		}
		if j, ok := p.next[s[i]]; ok {
			i, t = i+1, j
		} else {
			t = p.fail
		}
	}
	return -1, -1
}

// 搜索所有匹配的位置，[2]int成员分别是起始位置下标和模式串的下标，没有返回nil
func (this *AC) Find(s []byte) (o [][2]int) {
	t := 0
	for i, l := 0, len(s); i < l; {
		p := &this.State[t]
		if p.over >= 0 {
			o = append(o, [2]int{i - p.deep, p.over})
		}
		if j, ok := p.next[s[i]]; ok {
			i, t = i+1, j
		} else {
			t = p.fail
		}
	}
	return
}

// 为了方便存储Trie树，将之做成了list形式，所以必须使用本方法来提供新的节点
func (this *AC) newAcNode() int {
	n := this.Count
	this.State = append(this.State, AcNode{deep: -1, over: -1, fail: -1})
	this.Count++
	return n
}

// 以某个节点为根节点，向下生成Trie树，支持字符集匹配
func (this *AC) incStr(p, n int, s string) {
	st := &this.State[n]
	if s == "" {
		st.over = p
		return
	}
	head, tail := s[0], s[1:]
	if st.next == nil {
		st.next = make(map[byte]int)
	}
	i, ok := st.next[head]
	if !ok {
		i = this.newAcNode()         // 注意本行执行后，st指针很可能失效！
		this.State[n].next[head] = i // 所以要用回this.State[n]
	}
	this.incStr(p, i, tail)
}

// Brute-Force算法，即暴力搜索算法，最基本的字符串比对算法，首先对齐模式串和文本串的0位，然后
// 进行匹配，匹配失败则模式串向右移动一位，再次重新匹配，依次进行下去。
type StrBF struct {
	temp string
}

// 创建一个StrBF对象
func NewStrBF(s string) *StrBF {
	return &StrBF{s}
}

// 搜索第一个匹配的位置，没有返回-1
func (this StrBF) Index(s []byte) int {
	p := len(this.temp)
	q := len(s)
	for i := 0; i+p < q; i++ {
		if Compare(this.temp, s[i:]) {
			return i
		}
	}
	return -1
}

// 搜索所有匹配的位置，没有返回nil
func (this StrBF) Find(s []byte) (o []int) {
	p := len(this.temp)
	q := len(s)
	for i := 0; i+p < q; i++ {
		if Compare(this.temp, s[i:]) {
			o = append(o, i)
		}
	}
	return
}

// 固定位置匹配string格式的模式串和[]byte格式的文本串
func Compare(t string, s []byte) bool {
	for i, l := 0, len(t); i < l; i++ {
		if t[i] != s[i] {
			return false
		}
	}
	return true
}

// StrBM算法，即Boyer-Moore算法。这是一种目前常用的字符串匹配算法。horspool算法是其简化版。该
// 算法除了使用了horspool算法里用到的根据字符最后出现位置挪动模式串的原理外，还利用已经匹配的
// 后缀序列：如果该序列在模式串出现多次，则挪动模式串对齐倒数第二次出现的位置；如果后缀在模式串
// 中只出现末尾那一次，则使用模式串的前缀去匹配该后缀序列的后缀，取其最长匹配序列对齐；如果两者
// 没有非空的匹配序列，则模式串与当前文本串匹配右侧末端位置后面一个字符对齐。StrBM算法利用这两个规
// 则，选择移动距离最长的来挪动模式串，从而提升匹配速度。改进算法将horspool对齐规则替换为更好
// 的sunday规则（horspool和sunday规则基本一致，只是对齐指示字符不同）。
type StrBM struct {
	char [256]int
	next []int
	temp string
}

// 创建一个StrBM对象
// char数组记录了某个字符最后一次出现在模式串时，模式串到该字符的前缀的长度，否则为0。next切片
// 储存模式串后缀s[i:]对应的内部从右往左数第二个匹配以及其前端序列的总长度；如果不存在内部匹配
// 则储存与该后缀的某个后缀匹配的最长模式串前缀的长度；如果也不存在则设置为0。next切片最后一个
// 元素表示未匹配到后缀时（即匹配的后缀为空字符）的情况，即0。
func NewStrBM(s string) *StrBM {
	l := len(s)
	t := new(StrBM)
	t.temp = s
	for i := 0; i < l; i++ {
		t.char[int(s[i])] = i + 1
	}
	t.next = make([]int, l+1)
	for i := l - 1; i >= 0; i-- {
		j, d := i-1, l-i
		for ; j >= 0; j-- {
			if s[j:j+d] == s[i:] {
				break
			}
		}
		if j >= 0 {
			t.next[i] = j + d
			continue
		}
		for d--; d > 0; d-- {
			if s[l-d:] == s[:d] {
				break
			}
		}
		t.next[i] = d
	}
	return t
}

// 搜索第一个匹配的位置，没有返回-1
func (this *StrBM) Index(s []byte) int {
	var (
		p       = len(this.temp)
		q       = len(s)
		i, j, k int
	)
	for i = p - 1; i < q; {
		i, j, k = i+1, p-1, i
		for ; j >= 0; j, k = j-1, k-1 {
			if this.temp[j] != s[k] {
				break
			}
		}
		if j < 0 {
			return k + 1
		}
		if i < q {
			i += p - this.char[int(s[i])] // 采用sunday规则
		}
		k += p - this.next[j+1]
		if i < k {
			i = k
		}
	}
	return -1
}

// 搜索所有匹配的位置，没有返回nil
func (this *StrBM) Find(s []byte) (o []int) {
	var (
		p       = len(this.temp)
		q       = len(s)
		i, j, k int
	)
	for i = p - 1; i < q; {
		i, j, k = i+1, p-1, i
		for ; j >= 0; j, k = j-1, k-1 {
			if this.temp[j] != s[k] {
				break
			}
		}
		if j < 0 {
			o = append(o, k+1)
		}
		if i < q {
			i += p - this.char[int(s[i])] // 采用sunday规则
		}
		k += p - this.next[j+1]
		if i < k {
			i = k
		}
	}
	return
}

// Strhorspool算法的核心是利用字符在模式串中最后出现的位置来挪动模式串，模式串的移动是从左向右，
// 但匹配时却是从右向左。匹配过程中遇到不匹配字符时，移动模式串使文本串该字符与模式串中最后一个
// 该字符对齐，对齐最后一个的目的是防止漏解。如果模式串中不存在该字符，则移动模式串对齐该字符之
// 后的那个字符。算法关键是生成一个某字符在模式串最后出现位置的数组。注意根据该规则进行对齐可能
// 导致模式串反向移动，此时应将模式串强制右移一位。Strhorspool算法与sunday算法类似，区别在于用于
// 对齐模式串的指示字符不同。
type StrHorspool struct {
	char [256]int
	temp string
}

// 创建一个StrHorspool对象
// char数组记录了某个字符最后一次出现在模式串时，模式串到该字符的前缀的长度，否则为0
func NewStrHorspool(s string) *StrHorspool {
	t := new(StrHorspool)
	t.temp = s
	for i := 0; i < len(s); i++ {
		t.char[int(s[i])] = i + 1
	}
	return t
}

// 搜索第一个匹配的位置，没有返回-1
func (this *StrHorspool) Index(s []byte) int {
	var (
		p       = len(this.temp)
		q       = len(s)
		i, j, k int
	)
	for i = p - 1; i < q; {
		i, j, k = i+1, p-1, i
		for ; j >= 0; j, k = j-1, k-1 {
			if this.temp[j] != s[k] {
				break
			}
		}
		if j < 0 {
			return k + 1
		}
		if k += p - this.char[int(s[k])]; i < k {
			i = k
		}
	}
	return -1
}

// 搜索所有匹配的位置，没有返回nil
func (this *StrHorspool) Find(s []byte) (o []int) {
	var (
		p       = len(this.temp)
		q       = len(s)
		i, j, k int
	)
	for i = p - 1; i < q; {
		i, j, k = i+1, p-1, i
		for ; j >= 0; j, k = j-1, k-1 {
			if this.temp[j] != s[k] {
				break
			}
		}
		if j < 0 {
			o = append(o, k+1)
		}
		if k += p - this.char[int(s[k])]; i < k {
			i = k
		}
	}
	return
}

// StrKMP算法，全名Knuth-Morris-Pratt算法。其理论是，在进行字符串匹配时，如果遇到匹配失败的情
// 况，则根据已经匹配的部分挪动模式串的位置，使模式串的某个前缀与已匹配部分的某个后缀成功匹配，
// 并要求挪动距离最短即匹配长度最大，这是为了避免错失成功的匹配。注意，所谓“已匹配的部分”很明显
// 也是模式串的某个前缀。因此，StrKMP算法的关键是计算出模式串的某个前缀对应的最长的可作为其真后缀
// 的另一模式串前缀（显然也是该前缀的前缀），这里称其为最长双配序列。对模式串s建立一个整数数组，
// 称为n，n[i]的值为s[:i]的最长双配序列的长度。这一数组即被称为next数组。构建next数组的过程
// 和根据next数组进行匹配的过程是有异曲同工之妙的。
type StrKMP struct {
	next []int
	temp string
}

// 生成一个StrKMP对象，内部包含模式串和计算好的next数组。
func NewStrKMP(s string) *StrKMP {
	l := len(s)
	n := make([]int, l+1)

	// n[i]表示s长度为i的前缀s[:i]的最长双配序列的长度
	// 很显然n[0]、n[1]都是0，从n[2]开始算起
	for i, j := 2, 0; i <= l; {
		// 已知j=n[i-1]，亦即s[:i-1]的最长双配序列的长度
		// 则s[i-1]表示s[:i-1]后的第一个字符
		// s[j]表示s[:i-1]的最长双配序列后的第一个字符
		if s[i-1] == s[j] { // 说明s[:i]的最长双配序列的长度为j+1
			n[i] = j + 1    // 对n[i]赋值
			i, j = i+1, j+1 // 更新i、j，进入下一轮循环
		} else { // 字符不匹配，则需要寻找次长双配序列再次测试
			if j == 0 { // 空双配序列也不能满足延伸的要求s[i-1] == s[j]
				// n[i]为0，即默认值
				i++ // 注意j为0不需要赋值，进入下一轮循环
			} else { // 测试次长双配序列
				j = n[j] // 次长双配序列必然是最长双配序列的最长双配序列
			}
		}
	}
	return &StrKMP{n, s}
}

// 搜索第一个匹配的位置，没有返回-1
func (this *StrKMP) Index(s []byte) int {
	t := this.temp
	p := len(t)
	q := len(s)
	i, j := 0, 0
	for i+p < j+q {
		if s[i] == t[j] {
			i, j = i+1, j+1
			if j == p {
				return i - p
			}
		} else if j == 0 {
			i++
		} else {
			j = this.next[j]
		}
	}
	return -1
}

// 搜索所有匹配的位置，没有返回nil
func (this *StrKMP) Find(s []byte) (o []int) {
	t := this.temp
	p := len(t)
	q := len(s)
	i, j := 0, 0
	for i+p < j+q {
		if s[i] == t[j] {
			i, j = i+1, j+1
			if j == p {
				o = append(o, i-p)
				j = this.next[j]
			}
		} else {
			if j == 0 {
				i++
			} else {
				j = this.next[j]
			}
		}
	}
	return
}

// StrRK算法，即Robin-Karp算法，哈希检索算法。宗旨是对模式串求哈希ID，对匹配的文本也依次求哈希
// ID，在两个ID一致的进行逐字符比对的复核。该算法的关键是不能对每个哈希ID都要重新计算，这样算
// 法复杂度不会改变，而应采用由旧的ID根据新字符生成新的ID的哈希算法，比如说，累加求和，以及求
// 异或运算。复核是必须的，因为很可能存在哈希碰撞的情况。
type StrRK struct {
	Hash
	temp string
	sign int
}

// 用于StrRK算法的hash接口
type Hash interface {
	Feed(byte) // 吃进一个字符
	Free(byte) // 释放一个字符
	Reset()    // 重启hash缓存
	Sum() int  // 当前hash缓存
}

// 根据模式串和hash算法创建一个StrRK算法的对象
func NewStrRK(s string, h Hash) *StrRK {
	h.Reset()
	for i := 0; i < len(s); i++ {
		h.Feed(s[i])
	}
	return &StrRK{h, s, h.Sum()}
}

// 搜索第一个匹配的位置，没有返回-1
func (this *StrRK) Index(s []byte) int {
	p := len(this.temp)
	q := len(s)
	this.Reset()
	for i := 0; i < p; i++ {
		this.Feed(s[i])
	}
	// i,j分别表示生成下一个ID需要释放和接收的文本串字符的下标
	for i, j := 0, p; j < q; i, j = i+1, j+1 {
		if this.Sum() == this.sign {
			if Compare(this.temp, s[i:]) {
				return i
			}
		}
		this.Feed(s[j])
		this.Free(s[i])
	}
	return -1
}

// Find 搜索所有匹配的位置，没有返回nil
func (this *StrRK) Find(s []byte) (o []int) {
	p := len(this.temp)
	q := len(s)
	this.Reset()
	for i := 0; i < p; i++ {
		this.Feed(s[i])
	}
	// i,j分别表示生成下一个ID需要释放和接收的文本串字符的下标
	for i, j := 0, p; j < q; i, j = i+1, j+1 {
		if this.Sum() == this.sign {
			if Compare(this.temp, s[i:]) {
				o = append(o, i)
			}
		}
		this.Feed(s[j])
		this.Free(s[i])
	}
	return
}

// ExclusiveOr 利用异或运算实现了Hash接口
type ExclusiveOr int

func (this *ExclusiveOr) Feed(c byte) {
	*this ^= ExclusiveOr(c)
}

func (this *ExclusiveOr) Free(c byte) {
	*this ^= ExclusiveOr(c)
}

func (this *ExclusiveOr) Reset() {
	*this = 0
}

func (this *ExclusiveOr) Sum() int {
	return int(*this)
}

// PlusMinus 利用加减运算实现了Hash接口
type PlusMinus int

func (this *PlusMinus) Feed(c byte) {
	*this += PlusMinus(c)
}

func (this *PlusMinus) Free(c byte) {
	*this -= PlusMinus(c)
}

func (this *PlusMinus) Reset() {
	*this = 0
}

func (this *PlusMinus) Sum() int {
	return int(*this)
}

// StrSunday StrSunday算法的核心是利用字符在模式串中最后出现的位置来挪动模式串，模式串的移动是从左向右，但
// 匹配时却是从右向左。匹配过程中遇到不匹配字符时，以当前对齐位置下模式串末尾字符对应的文本串字
// 符的下一个字符作为标准，移动模式串使模式串中最后一个该字符与文本串该字符对齐，对齐最后一个的
// 目的是防止漏解。如果模式串中不存在该字符，则移动模式串对齐该字符之后的那个位置。算法关键是生
// 成一个某字符在模式串最后出现位置的数组。因为以模式串末尾后的字符为标准对齐，不会出现模式串退
// 步的情况。StrSunday算法与horspool算法类似，区别在于用于对齐模式串的指示字符不同。
type StrSunday struct {
	char [256]int
	temp string
}

// NewStrSunday 创建一个StrSunday对象
// char数组记录了某个字符最后一次出现在模式串时，模式串到该字符的前缀的长度，否则为0
func NewStrSunday(s string) *StrSunday {
	t := new(StrSunday)
	t.temp = s
	for i := 0; i < len(s); i++ {
		t.char[int(s[i])] = i + 1
	}
	return t
}

// Index 搜索第一个匹配/**/的位置，没有返回-1
func (this *StrSunday) Index(s []byte) int {
	var (
		p       = len(this.temp)
		q       = len(s)
		i, j, k int
	)
	for i = p - 1; i < q; {
		i, j, k = i+1, p-1, i
		for ; j >= 0; j, k = j-1, k-1 {
			if this.temp[j] != s[k] {
				break
			}
		}
		if j < 0 {
			return k + 1
		}
		if i < q {
			i = i + p - this.char[int(s[i])]
		}
	}
	return -1
}

// Find 搜索所有匹配的位置，没有返回nil
func (this *StrSunday) Find(s []byte) (o []int) {
	var (
		p       = len(this.temp)
		q       = len(s)
		i, j, k int
	)
	for i = p - 1; i < q; {
		i, j, k = i+1, p-1, i
		for ; j >= 0; j, k = j-1, k-1 {
			if this.temp[j] != s[k] {
				break
			}
		}
		if j < 0 {
			o = append(o, k+1)
		}
		if i < q {
			i = i + p - this.char[int(s[i])]
		}
	}
	return
}

//SundayStrStr SundayStrStr
func SundayStrStr(haystack, needle string) int {
	if needle == "" { //排除特殊情况
		return 0
	}
	m := make(map[byte]int, len(needle))
	for i := 0; i < len(needle); i++ {
		m[needle[i]] = i //记录模式串每一个字符的位置(重复取最右边)
	}
	j := 0 //模式串指针
	for i := 0; i < len(haystack); i++ {
		if haystack[i] != needle[j] {
			//如果发生不匹配,i跳到主串对应模式串长度下一个字符
			i = len(needle) - j + i
			if i >= len(haystack) { //i越界haystack说明找不到
				return -1
			}
			v, ok := m[haystack[i]]
			if ok { //字符存在于模式串中,则用该字符对齐主串与模式串,从模式串头开始比较
				j = 0
				i = i - v - 1
				continue
			} else { //字符不存在于模式串中,模式串直接滑动到该字符的下一位从模式串头比较
				j = 0
				continue
			}
		}
		if j == len(needle)-1 {
			return i + 1 - len(needle)
		}
		j++
	}
	return -1
}

// KMPStrStr KMPStrStr
func KMPStrStr(haystack, needle string) int {
	//KMP改进算法   模式串连续重复字符超过2个,如"aaaaaag"
	n, m := len(haystack), len(needle)
	//排除 needle=""
	if m == 0 {
		return 0
	}
	//next数组表示模式串当前位(从0数)的最长公共前后缀
	//nextval是对next的改进,模式串连续重复字符超过2个的情况
	nextval := make([]int, m)
	for i, j := 1, 0; i < m; i++ {
		for j > 0 && needle[i] != needle[j] {
			j = nextval[j-1]
		}
		if needle[i] == needle[j] {
			j++
		}
		nextval[i] = j
		//如果与上一个字符相同b并且与下一个字符相同,则nextval与上一个字符相等
		if i < m-1 && needle[i] == needle[i-1] && needle[i] == needle[i+1] {
			nextval[i] = nextval[i-1]
		}
	}
	//遍历haystack
	for i, j := 0, 0; i < n; i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = nextval[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == m {
			return i - m + 1
		}
	}
	return -1
}

//BoomStrStr BoomStrStr
func BoomStrStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i] == needle[0] {
			if haystack[i:i+len(needle)] == needle {
				return i
			}
		}
	}
	return -1
}
