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
