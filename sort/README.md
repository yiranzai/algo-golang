# 排序

排序相关

## 目录

---

<!--ts-->
   * [排序](README.md#排序)
      * [目录](README.md#目录)
      * [<a href="https://zh.wikipedia.org/wiki/%E6%8E%92%E5%BA%8F%E7%AE%97%E6%B3%95" rel="nofollow">排序算法</a>](README.md#排序算法)
         * [<a href="https://zh.wikipedia.org/wiki/%E5%86%92%E6%B3%A1%E6%8E%92%E5%BA%8F" rel="nofollow">冒泡排序</a>](README.md#冒泡排序)
         * [<a href="https://zh.wikipedia.org/wiki/%E5%BF%AB%E9%80%9F%E6%8E%92%E5%BA%8F" rel="nofollow">快速排序</a>](README.md#快速排序)
            * [递归原地快排](README.md#递归原地快排)
            * [循环原地快排](README.md#循环原地快排)
            * [递归三切快排](README.md#递归三切快排)
            * [循环三切快排](README.md#循环三切快排)
         * [<a href="https://zh.wikipedia.org/wiki/%E9%80%89%E6%8B%A9%E6%8E%92%E5%BA%8F" rel="nofollow">选择排序</a>](README.md#选择排序)
         * [<a href="https://zh.wikipedia.org/wiki/%E6%8F%92%E5%85%A5%E6%8E%92%E5%BA%8F" rel="nofollow">插入排序</a>](README.md#插入排序)
         * [<a href="https://zh.wikipedia.org/wiki/%E5%A0%86%E6%8E%92%E5%BA%8F" rel="nofollow">堆排序</a>](README.md#堆排序)
         * [<a href="https://zh.wikipedia.org/wiki/%E5%BD%92%E5%B9%B6%E6%8E%92%E5%BA%8F" rel="nofollow">归并排序</a>](README.md#归并排序)
         * [<a href="https://zh.wikipedia.org/wiki/%E5%B8%8C%E5%B0%94%E6%8E%92%E5%BA%8F" rel="nofollow">希尔排序</a>](README.md#希尔排序)
         * [<a href="https://zh.wikipedia.org/wiki/%E8%AE%A1%E6%95%B0%E6%8E%92%E5%BA%8F" rel="nofollow">计数排序</a>](README.md#计数排序)
         * [<a href="https://zh.wikipedia.org/wiki/%E6%A1%B6%E6%8E%92%E5%BA%8F" rel="nofollow">桶排序</a>](README.md#桶排序)
         * [<a href="https://zh.wikipedia.org/wiki/%E5%9F%BA%E6%95%B0%E6%8E%92%E5%BA%8F" rel="nofollow">基数排序</a>](README.md#基数排序)

<!-- Added by: runner, at: Tue May 24 04:54:03 UTC 2022 -->

<!--te-->

---

## [排序算法](https://zh.wikipedia.org/wiki/排序算法)

在计算机科学与数学中，一个排序算法（英语：Sorting
algorithm）是一种能将一串资料依照特定排序方式进行排列的一种算法。最常用到的排序方式是数值顺序以及字典顺序。有效的排序算法在一些算法（例如搜索算法与合并算法）中是重要的，如此这些算法才能得到正确解答。排序算法也用在处理文字资料以及产生人类可读的输出结果。基本上，排序算法的输出必须遵守下列两个原则：

输出结果为递增序列（递增是针对所需的排序顺序而言） 输出结果是原输入的一种排列、或是重组
虽然排序算法是一个简单的问题，但是从计算机科学发展以来，在此问题上已经有大量的研究。举例而言，冒泡排序在1956年就已经被研究。虽然大部分人认为这是一个已经被解决的问题，有用的新算法仍在不断的被发明。

### [冒泡排序](https://zh.wikipedia.org/wiki/冒泡排序)

从无序区透过交换找出最大元素放到有序区前端。

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

### [快速排序](https://zh.wikipedia.org/wiki/快速排序)

> “茴“字的三种写法

在区间中随机挑选一个元素作基准，将小于基准的元素放在基准之前，大于基准的元素放在基准之后，再分别对小数区与大数区进行排序。

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

### [选择排序](https://zh.wikipedia.org/wiki/选择排序)

在无序区里找一个最小的元素跟在有序区的后面。对数组：比较得多，换得少。

```go
package main
```

### [插入排序](https://zh.wikipedia.org/wiki/插入排序)

把无序区的第一个元素插入到有序区的合适的位置。对数组：比较得少，换得多。

```go
package main
```

### [堆排序](https://zh.wikipedia.org/wiki/堆排序)

从堆顶把根卸出来放在有序区之前，再恢复堆。

```go
package main
```

### [归并排序](https://zh.wikipedia.org/wiki/归并排序)

把数据分为两段，从两段中逐个选最小的元素移入新数据段的末尾。可从上到下或从下到上进行。

```go
package main
```

### [希尔排序](https://zh.wikipedia.org/wiki/希尔排序)

每一轮按照事先决定的间隔进行插入排序，间隔会依次缩小，最后一次一定要是1。

```go
package main
```

### [计数排序](https://zh.wikipedia.org/wiki/计数排序)

统计小于等于该元素值的元素的个数i，于是该元素就放在目标数组的索引i位（i≥0）。

```go
package main
```

### [桶排序](https://zh.wikipedia.org/wiki/桶排序)

将值为i的元素放入i号桶，最后依次把桶里的元素倒出来。

```go
package main
```

### [基数排序](https://zh.wikipedia.org/wiki/基数排序)

一种多关键字的排序算法，可用桶排序实现。

```go
package main
```

