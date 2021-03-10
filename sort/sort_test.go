package sort

import "testing"

var (
	nums   = []int{70, 90, 30, 50, 70, 81, 89, 23, 34, 39, 57, 7, 54, 65, 10, 40, 80}
	result = []int{7, 10, 23, 30, 34, 39, 40, 50, 54, 57, 65, 70, 70, 80, 81, 89, 90}
)

func TestBubbleSort(t *testing.T) {
	eqResult(t, BubbleSort(nums))
}

func TestQuickSortDeep(t *testing.T) {
	QuickSortDeep(&nums)
	eqResult(t, nums)
}

func TestQuickSortLoop(t *testing.T) {
	QuickSortLoop(&nums)
	eqResult(t, nums)
}

func eqResult(t *testing.T, nums []int) {
	if len(nums) != len(result) {
		t.Error(nums)
	}
	for key, value := range nums {
		if value != result[key] {
			t.Error(key, value, result[key])
		}
	}
}
