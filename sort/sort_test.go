package sort

import (
	"testing"

	"gotest.tools/v3/assert"
)

var (
	nums = [][]int{{70, 90, 30, 50, 70, 81, 89, 23, 34, 39, 57, 7, 54, 65, 10, 40, 80},
		{80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80,
			80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80}}
	result = [][]int{{7, 10, 23, 30, 34, 39, 40, 50, 54, 57, 65, 70, 70, 80, 81, 89, 90},
		{80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80,
			80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80}}
)

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, num := range nums {
			BubbleSort(num)
		}
	}
}

func TestBubbleSort(t *testing.T) {
	for i, num := range nums {
		assert.DeepEqual(t, BubbleSort(num), result[i])
	}

}

func BenchmarkQuickSortDeep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, num := range nums {
			QuickSortDeep(num)
		}
	}

}

func TestQuickSortDeep(t *testing.T) {
	for i, num := range nums {
		QuickSortDeep(num)
		assert.DeepEqual(t, num, result[i])
	}

}

func BenchmarkQuickSortDeep3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, num := range nums {
			QuickSortDeep3(num)
		}
	}

}

func TestQuickSortDeep3(t *testing.T) {
	for i, num := range nums {
		QuickSortDeep3(num)
		assert.DeepEqual(t, num, result[i])
	}

}

func BenchmarkQuickSortLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, num := range nums {
			QuickSortLoop(num)
		}
	}

}

func TestQuickSortLoop(t *testing.T) {
	for i, num := range nums {
		QuickSortLoop(num)
		assert.DeepEqual(t, num, result[i])
	}
}

func BenchmarkQuickSortLoop2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, num := range nums {
			QuickSortLoop2(num)
		}
	}
}

func TestQuickSortLoop2(t *testing.T) {
	for i, num := range nums {
		QuickSortLoop2(num)
		assert.DeepEqual(t, num, result[i])
	}
}
