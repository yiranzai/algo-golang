package offer

import (
	"github.com/yiranzai/go-utils/math"
)

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
// https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/
func findNumberIn2DArray(matrix [][]int, target int) bool {
	yLen := len(matrix)
	if yLen == 0 {
		return false
	}
	xLen := len(matrix[0])
	if xLen == 0 {
		return false
	}
	if matrix[0][0] > target {
		return false
	}
	yMid := yLen >> 1
	xMid := xLen >> 1

	if matrix[yMid][xMid] > target {

	}

	return false
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type Node struct {
	Val  int
	Prev *Node
	Next *Node
}

/**
 * https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
	var stack []int
	if head == nil {
		return stack
	}
	stack = append(stack, head.Val)
	for head.Next != nil {
		head = head.Next
		stack = append(stack, head.Val)
	}

	length := len(stack)

	for i := 0; i < length>>1; i++ {
		stack[i], stack[length-1-i] = stack[length-1-i], stack[i]
	}
	return stack
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof/
 * https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 *
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	var root *TreeNode

	return root
}

/**
*  https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/  -------------------------------------
 */

type CQueue struct {
	Current, Head *Node
}

func Constructor() CQueue {
	return CQueue{}
}

func (c *CQueue) AppendTail(value int) {
	node := &Node{
		Val:  value,
		Prev: nil,
		Next: c.Current,
	}
	if c.Current != nil {
		c.Current.Prev = node
	}
	c.Current = node
	if c.Head == nil {
		c.Head = node
	}
}

func (c *CQueue) DeleteHead() int {
	if c.Head == nil {
		return -1
	}
	res := c.Head.Val
	c.Head = c.Head.Prev
	if c.Head != nil {
		c.Head.Next = nil
	}
	return res
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

/**
*  https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/  -------------------------------------
 */

// https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof/
var bimap = make(map[int]int)

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	a := bimap[n-1]
	if a == 0 {
		a = fib(n - 1)
		bimap[n-1] = a
	}

	b := bimap[n-2]
	if b == 0 {
		b = fib(n - 2)
		bimap[n-2] = b
	}

	return (a + b) % 1000000007
}

// https://leetcode-cn.com/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/
// https://leetcode-cn.com/problems/climbing-stairs/
func numWays(n int) int {
	var result int
	return result
}

// https://leetcode-cn.com/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/
// https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/

func minArray(numbers []int) int {
	return findMin(numbers, 0, len(numbers)-1)
}

func findMin(nums []int, start, end int) int {
	if start == end {
		return nums[end]
	}
	// 开头比结尾还小，直接返回开头的值
	if nums[start] < nums[end] {
		return nums[start]
	}
	mid := start + ((end - start) >> 1)

	return math.MinInt(findMin(nums, start, mid), findMin(nums, mid+1, end))
}

// https://leetcode-cn.com/problems/ju-zhen-zhong-de-lu-jing-lcof/
// https://leetcode-cn.com/problems/word-search/
var (
	lengthX int
	lengthY int
	lengthW int
)

func exist(board [][]byte, word string) bool {
	lengthW = len(word)
	if lengthW == 0 {
		return true
	}
	lengthY = len(board)
	if lengthY != 0 {
		lengthX = len(board[0])
	}
	w := []byte(word)
	for y := 0; y < lengthY; y++ {
		for x := 0; x < lengthX; x++ {
			if w[0] == board[y][x] {
				if searchString(board, w[1:], x, y) {
					return true
				} else {
					// 没找到还原这个字符
					board[y][x] = w[0]
				}
			}
		}
	}

	return false
}

func searchString(board [][]byte, word []byte, x, y int) bool {
	if len(word) == 0 {
		return true
	}
	board[y][x] = ' '
	// 左
	if x != 0 {
		if board[y][x-1] == word[0] {
			if searchString(board, word[1:], x-1, y) {
				return true
			} else {
				board[y][x-1] = word[0]
			}
		}
	}

	// 右
	if x != lengthX-1 {
		if board[y][x+1] == word[0] {
			if searchString(board, word[1:], x+1, y) {
				return true
			} else {
				board[y][x+1] = word[0]
			}
		}
	}

	// 上
	if y != 0 {
		if board[y-1][x] == word[0] {
			if searchString(board, word[1:], x, y-1) {
				return true
			} else {
				board[y-1][x] = word[0]
			}
		}
	}

	// 下
	if y != lengthY-1 {
		if board[y+1][x] == word[0] {
			if searchString(board, word[1:], x, y+1) {
				return true
			} else {
				board[y+1][x] = word[0]
			}
		}
	}
	return false
}

// https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/
func movingCount(m int, n int, k int) int {
	var baseY, count int
	var res [][]bool
	res = make([][]bool, m)
	count = 0
	for y := 0; y < m; y++ {
		baseY = sumByte(y)
		if baseY > k {
			break
		}
		res[y] = make([]bool, n)
		for x := 0; x < n; x++ {
			// 当上边或者左边已经遍历过且是可达的,或者处于起点
			if (x > 0 && res[y][x-1] == true) || (y > 0 && res[y-1][x] == true) || (x == 0 && y == 0) {
				if sumByte(x) <= k-baseY {
					res[y][x] = true
					count++
				}
			}
		}
	}
	return count
}

//sumByte 统计数位
func sumByte(num int) int {
	if num == 100 {
		return 1
	}
	if num < 10 {
		return num
	}
	return num%10 + num/10
}

// https://leetcode-cn.com/problems/jian-sheng-zi-lcof/
// https://leetcode-cn.com/problems/integer-break/
func cuttingRope(n int) int {
	return n
}
