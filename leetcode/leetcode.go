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

// https://leetcode-cn.com/problems/lru-cache/
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

// https://leetcode-cn.com/problems/lfu-cache/
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
