package go_algorithm_pattern

//TreeNode 树的节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//Node 链表节点
type Node struct {
	Val  int
	Next *TreeNode
}

func BinarySearch(nums []int, target int) int {
	length := len(nums)
	left := 0
	right := length

	for left < right {
		// 避免数据超范围
		mid := left + ((right - left) >> 1)
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		} else {
			return mid
		}
	}
	return right
}

func Kmp(pat, txt string) {

}
