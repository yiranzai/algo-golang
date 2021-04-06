# 排序

排序相关

## 目录

---

<!--ts-->

* [排序](#排序)
    * [目录](#目录)

<!-- Added by: runner, at: Wed Mar 31 15:24:40 UTC 2021 -->

<!--te-->

---

## 排序算法

### 冒泡排序

```go
package main

func BubbleSort(nums []int) []int {
	num := len(nums)
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}
```

### 快速排序

“茴“字的三种写法

#### 递归原地快排

```go
package main

func QuickSortDeep(nums []int) {
	quickSortDeep(nums, 0, len(nums)-1)
}
func quickSortDeep(nums []int, start, end int) {

	num := len(nums)
	if num <= 1 {
		return
	}

	if start >= end {
		return
	}

	value := nums[end]

	var j = start
	for i := start; i < end; i++ {
		if nums[i] < value {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	nums[j], nums[end] = nums[end], nums[j]
	quickSortDeep(nums, start, j-1)
	quickSortDeep(nums, j+1, end)
}
```

#### 循环原地快排

```go
package main

//QuickSortLoop 原地快排非递归
func QuickSortLoop(nums []int) {
	length := len(nums) - 1
	if length == 0 {
		return
	}
	temp := make([][]int, 0, length)
	temp = append(temp, []int{0, length})
	for len(temp) > 0 {
		pool := temp
		temp = make([][]int, 0, length)
		for _, ints := range pool {
			start, end := ints[0], ints[1]
			if start >= end {
				continue
			}
			endValue := nums[end]
			var j = start
			for i := start; i < end; i++ {
				if nums[i] < endValue {
					nums[i], nums[j] = nums[j], nums[i]
					j++
				}
			}
			nums[j], nums[end] = nums[end], nums[j]

			temp = append(temp, []int{start, j - 1})
			temp = append(temp, []int{j + 1, end})
		}
	}
}
```

#### 递归三切快排

```go
package main

func QuickSortDeep3(nums []int) {
	quickSortDeep3(nums, 0, len(nums)-1)
}

func quickSortDeep3(nums []int, start, end int) {
	num := len(nums)
	if num <= 1 {
		return
	}

	if start >= end {
		return
	}

	value := nums[end]

	var j = start
	var m = start
	for i := start; i < end; i++ {
		if nums[i] > value {
			continue
		}
		if nums[i] < value {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
		m++
	}
	nums[j], nums[end] = nums[end], nums[j]
	quickSortDeep3(nums, start, j-1)
	quickSortDeep3(nums, m+1, end)
}

```

#### 循环三切快排

```go
package main

func QuickSortLoop2(nums []int) {
	length := len(nums) - 1
	if length == 0 {
		return
	}
	temp := make([][]int, 0, length)
	temp = append(temp, []int{0, length})
	for len(temp) > 0 {
		pool := temp
		temp = make([][]int, 0, length)
		for _, ints := range pool {
			start, end := ints[0], ints[1]
			if start >= end {
				continue
			}
			endValue := nums[end]
			var j = start
			var m = start
			for i := start; i < end; i++ {
				if nums[i] > endValue {
					continue
				}
				if nums[i] < endValue {
					nums[i], nums[j] = nums[j], nums[i]
					j++
				}
				m++
			}
			nums[j], nums[end] = nums[end], nums[j]

			temp = append(temp, []int{start, j - 1})
			temp = append(temp, []int{m + 1, end})
		}
	}
}
```
