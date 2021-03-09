package offer

import (
	"gotest.tools/v3/assert"
	"testing"
)

func TestFindRepeatNumber(t *testing.T) {
	nums := []int{0, 2, 3, 1, 0, 2, 5, 3}
	assert.Equal(t, FindRepeatNumber(nums), 0)
	nums = []int{2, 3, 1, 0, 2, 5, 3}
	assert.Equal(t, FindRepeatNumber(nums), 2)
	nums = []int{0, 2, 3, 1, 2, 5, 3}
	assert.Equal(t, FindRepeatNumber(nums), 2)

	nums = []int{0, 2, 3, 1, 0, 2, 5, 3}
	assert.Equal(t, FindRepeatNumber2(nums), 0)
	nums = []int{2, 3, 1, 0, 2, 5, 3}
	assert.Equal(t, FindRepeatNumber2(nums), 2)
	nums = []int{0, 2, 3, 1, 2, 5, 3}
	assert.Equal(t, FindRepeatNumber2(nums), 2)
}

func TestFindNumberIn2DArray(t *testing.T) {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	assert.Equal(t, findNumberIn2DArray(matrix, 5), true)
	assert.Equal(t, findNumberIn2DArray(matrix, 20), false)
}

func Test_reversePrint(t *testing.T) {
	headList := []int{1, 3, 2}
	head := &ListNode{Val: 4}
	b := head
	for _, v := range headList {
		head.Next = &ListNode{Val: v}
		head = head.Next
	}
	assert.DeepEqual(t, reversePrint(b), []int{2, 3, 1, 4})

	headList = []int{1, 3, 2, 7}
	head = &ListNode{Val: 4}
	b = head
	for _, v := range headList {
		head.Next = &ListNode{Val: v}
		head = head.Next
	}
	assert.DeepEqual(t, reversePrint(b), []int{7, 2, 3, 1, 4})
}

func Test_CQueue(t *testing.T) {
	c := Constructor()

	assert.Equal(t, c.DeleteHead(), -1)
	assert.Equal(t, c.DeleteHead(), -1)

	c.AppendTail(5)
	c.AppendTail(2)

	assert.Equal(t, c.DeleteHead(), 5)

	assert.Equal(t, c.DeleteHead(), 2)
	assert.Equal(t, c.DeleteHead(), -1)
}

func Test_fib(t *testing.T) {
	assert.Equal(t, fib(5), 5)
	assert.Equal(t, fib(2), 1)
	assert.Equal(t, fib(3), 2)
	assert.Equal(t, fib(4), 3)
}
