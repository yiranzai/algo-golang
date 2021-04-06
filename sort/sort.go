package sort

//BubbleSort 冒泡排序
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

//QuickSortDeep 原地快排递归
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

//QuickSortLoop2 三切原地快排非递归
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

// https://www.jianshu.com/p/d70aeccaee19
//QuickSortDeep 三向切分快速排序
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
