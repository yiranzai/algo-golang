# Leetcode

## 目录

---

<!--ts-->

* [Leetcode](#leetcode)
    * [目录](#目录)
        * [<a href="https://leetcode-cn.com/problems/minimum-path-sum/" rel="nofollow">64. 最小路径和</a>](#64-最小路径和)
            * [Solution](#solution)
        * [<a href="https://leetcode-cn.com/problems/unique-paths/" rel="nofollow">62. 不同路径</a>](#62-不同路径)
            * [Solution](#solution-1)
        * [<a href="https://leetcode-cn.com/problems/unique-paths-ii/" rel="nofollow">63. 不同路径 II</a>](#63-不同路径-ii)
            * [Solution](#solution-2)
        * [<a href="https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/" rel="nofollow">82. 删除排序链表中的重复元素 II</a>](#82-删除排序链表中的重复元素-ii)
            * [Solution](#solution-3)
        * [<a href="https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/" rel="nofollow">83. 删除排序链表中的重复元素</a>](#83-删除排序链表中的重复元素)
            * [Solution](#solution-4)

<!-- Added by: runner, at: Thu Mar 25 16:28:31 UTC 2021 -->

<!--te-->

---

### [64\. 最小路径和](https://leetcode-cn.com/problems/minimum-path-sum/)

Difficulty: **中等**

给定一个包含非负整数的 `_m_ x _n_`  网格  `grid` ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

**说明：**每次只能向下或者向右移动一步。

**示例 1：**

![](https://assets.leetcode.com/uploads/2020/11/05/minpath.jpg)

```
输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
输出：7
解释：因为路径 1→3→1→1→1 的总和最小。
```

**示例 2：**

```
输入：grid = [[1,2,3],[4,5,6]]
输出：12
```

**提示：**

- `m == grid.length`
- `n == grid[i].length`
- `1 <= m, n <= 200`
- `0 <= grid[i][j] <= 100`

#### Solution

Language: **GO**

```go
​func minPathSum(grid [][]int) int {
lengY := len(grid)-1
lengX := len(grid[0])-1
for y := 0; y <= lengY; y++ {
for x := 1; x <= lengX; x++ {
if y == 0 {
grid[y][x] += grid[y][x-1]
continue
}
grid[y][x] += MinInt(grid[y-1][x], grid[y][x-1])
}
if y == lengY {
continue
}
grid[y+1][0] += grid[y][0]
}

return grid[lengY][lengX]
}

func MinInt(vars ...int) int {
min := vars[0]

for _, i := range vars {
if min > i {
min = i
}
}

return min
}
```

### [62\. 不同路径](https://leetcode-cn.com/problems/unique-paths/)

Difficulty: **中等**

一个机器人位于一个 `m x n`网格的左上角 （起始点在下图中标记为 “Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。

问总共有多少条不同的路径？

**示例 1：**

![](https://assets.leetcode.com/uploads/2018/10/22/robot_maze.png)

```
输入：m = 3, n = 7
输出：28
```

**示例 2：**

```
输入：m = 3, n = 2
输出：3
解释：
从左上角开始，总共有 3 条路径可以到达右下角。
1\. 向右 -> 向下 -> 向下
2\. 向下 -> 向下 -> 向右
3\. 向下 -> 向右 -> 向下
```

**示例 3：**

```
输入：m = 7, n = 3
输出：28
```

**示例 4：**

```
输入：m = 3, n = 3
输出：6
```

**提示：**

- `1 <= m, n <= 100`
- 题目数据保证答案小于等于 `2 * 10<sup>9</sup>`

#### Solution

Language: **GO**

```go
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
```

### [63\. 不同路径 II](https://leetcode-cn.com/problems/unique-paths-ii/)

Difficulty: **中等**

一个机器人位于一个 _m x n_ 网格的左上角 （起始点在下图中标记为“Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。

现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？

![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/22/robot_maze.png)

网格中的障碍物和空位置分别用 `1` 和 `0` 来表示。

**示例 1：**

![](https://assets.leetcode.com/uploads/2020/11/04/robot1.jpg)

```
输入：obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
输出：2
解释：
3x3 网格的正中间有一个障碍物。
从左上角到右下角一共有 2 条不同的路径：
1\. 向右 -> 向右 -> 向下 -> 向下
2\. 向下 -> 向下 -> 向右 -> 向右
```

**示例 2：**

![](https://assets.leetcode.com/uploads/2020/11/04/robot2.jpg)

```
输入：obstacleGrid = [[0,1],[0,0]]
输出：1
```

**提示：**

- `m == obstacleGrid.length`
- `n == obstacleGrid[i].length`
- `1 <= m, n <= 100`
- `obstacleGrid[i][j]` 为 `0` 或 `1`

#### Solution

Language: **GO**

```golang
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
}else {
dp[x] = 1
}
continue
}
dp[x] += dp[x-1]
}
if boolY == lenY + 1 {
return 0
}
}

return dp[lenX]
}

```

### [82\. 删除排序链表中的重复元素 II](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/)

Difficulty: **中等**

存在一个按升序排列的链表，给你这个链表的头节点 `head` ，请你删除链表中所有存在数字重复情况的节点，只保留原始链表中 **没有重复出现**的数字。

返回同样按升序排列的结果链表。

**示例 1：**

![](https://assets.leetcode.com/uploads/2021/01/04/linkedlist1.jpg)

```
输入：head = [1,2,3,3,4,4,5]
输出：[1,2,5]
```

**示例 2：**

![](https://assets.leetcode.com/uploads/2021/01/04/linkedlist2.jpg)

```
输入：head = [1,1,1,2,3]
输出：[2,3]
```

**提示：**

* 链表中节点数目在范围 `[0, 300]` 内
* `-100 <= Node.val <= 100`
* 题目数据保证链表已经按升序排列

#### Solution

Language: **GO**

```
​/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
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
```

### [83\. 删除排序链表中的重复元素](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/)

Difficulty: **简单**

存在一个按升序排列的链表，给你这个链表的头节点 `head` ，请你删除所有重复的元素，使每个元素 **只出现一次** 。

返回同样按升序排列的结果链表。

**示例 1：**

![](https://assets.leetcode.com/uploads/2021/01/04/list1.jpg)

```
输入：head = [1,1,2]
输出：[1,2]
```

**示例 2：**

![](https://assets.leetcode.com/uploads/2021/01/04/list2.jpg)

```
输入：head = [1,1,2,3,3]
输出：[1,2,3]
```

**提示：**

* 链表中节点数目在范围 `[0, 300]` 内
* `-100 <= Node.val <= 100`
* 题目数据保证链表已经按升序排列

#### Solution

Language: **GO**

```
​/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
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
```

### [70\. 爬楼梯](https://leetcode-cn.com/problems/climbing-stairs/)

Difficulty: **简单**

假设你正在爬楼梯。需要 _n_ 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

**注意：**给定 _n_ 是一个正整数。

**示例 1：**

```
输入： 2
输出： 2
解释： 有两种方法可以爬到楼顶。
1\.  1 阶 + 1 阶
2\.  2 阶
```

**示例 2：**

```
输入： 3
输出： 3
解释： 有三种方法可以爬到楼顶。
1\.  1 阶 + 1 阶 + 1 阶
2\.  1 阶 + 2 阶
3\.  2 阶 + 1 阶
```

#### Solution

Language: ****

```
​func climbStairs(n int) int {
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

```

### [55\. 跳跃游戏](https://leetcode-cn.com/problems/jump-game/)

Difficulty: **中等**

给定一个非负整数数组 `nums` ，你最初位于数组的 **第一个下标** 。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

判断你是否能够到达最后一个下标。

**示例 1：**

```
输入：nums = [2,3,1,1,4]
输出：true
解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
```

**示例 2：**

```
输入：nums = [3,2,1,0,4]
输出：false
解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。
```

**提示：**

* `1 <= nums.length <= 3 * 10<sup>4</sup>`
* `0 <= nums[i] <= 10<sup>5</sup>`

#### Solution

Language: ****

```
​func canJump(nums []int) bool {
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

```

### [45\. 跳跃游戏 II](https://leetcode-cn.com/problems/jump-game-ii/)

Difficulty: **中等**

给定一个非负整数数组，你最初位于数组的第一个位置。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

你的目标是使用最少的跳跃次数到达数组的最后一个位置。

**示例:**

```
输入: [2,3,1,1,4]
输出: 2
解释: 跳到最后一个位置的最小跳跃数是 2。
     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
```

**说明:**

假设你总是可以到达数组的最后一个位置。

#### Solution

Language: ****

```
​func jump(nums []int) int {
	length := len(nums) - 1
	if length == 0 {
		return 0
	}
	dps := make([]int, length+1)

	for i := 0; i <= length; i++ {
		for j := 1; j <= nums[i] ; j++ {
			if dps[i] == 0 || dps[i+j] == 0 || dps[i+j] > dps[i]{
				dps[i+j] = dps[i] + 1
			}

			if i+j >= length {
				return dps[length]
			}
		}
	}

	return dps[length]
}
```

### [139\. 单词拆分](https://leetcode-cn.com/problems/word-break/)

Difficulty: **中等**

给定一个**非空**字符串 _s_ 和一个包含**非空**单词的列表 _wordDict_，判定 _s_ 是否可以被空格拆分为一个或多个在字典中出现的单词。

**说明：**

* 拆分时可以重复使用字典中的单词。
* 你可以假设字典中没有重复的单词。

**示例 1：**

```
输入: s = "leetcode", wordDict = ["leet", "code"]
输出: true
解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
```

**示例 2：**

```
输入: s = "applepenapple", wordDict = ["apple", "pen"]
输出: true
解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
     注意你可以重复使用字典中的单词。
```

**示例 3：**

```
输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
输出: false
```

#### Solution

Language: ****

```
​var (
	dict map[string]bool
	dp   map[int]int
	maxLength int
)

func wordBreak(s string, wordDict []string) bool {
	dp = make(map[int]int)
	dict = make(map[string]bool)
	for _, s2 := range wordDict {
		maxLength = MaxInt(len(s2), maxLength)
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

func MaxInt(vars ...int) int {
	max := vars[0]

	for _, i := range vars {
		if max < i {
			max = i
		}
	}

	return max
}
```

### [1143\. 最长公共子序列](https://leetcode-cn.com/problems/longest-common-subsequence/)

Difficulty: **中等**

给定两个字符串 `text1` 和 `text2`，返回这两个字符串的最长公共子序列的长度。

一个字符串的 _子序列 _是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。  
例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。两个字符串的「公共子序列」是这两个字符串所共同拥有的子序列。

若这两个字符串没有公共子序列，则返回 0。

**示例 1:**

```
输入：text1 = "abcde", text2 = "ace" 
输出：3  
解释：最长公共子序列是 "ace"，它的长度为 3。
```

**示例 2:**

```
输入：text1 = "abc", text2 = "abc"
输出：3
解释：最长公共子序列是 "abc"，它的长度为 3。
```

**示例 3:**

```
输入：text1 = "abc", text2 = "def"
输出：0
解释：两个字符串没有公共子序列，返回 0。
```

**提示:**

* `1 <= text1.length <= 1000`
* `1 <= text2.length <= 1000`
* 输入的字符串只含有小写英文字符。

#### Solution

Language: ****

```
​func longestCommonSubsequence(text1 string, text2 string) int {
	var t1Len, t2Len int
	t1Len = len(text1)
	t2Len = len(text2)
	dp := make([][]int, t1Len+1)
	for i := 0; i <= 1; i++ {
		dp[i] = make([]int, t2Len+1)
	}
	k,l := 0,1
	for i := 1; i <= t1Len; i++ {
		for j := 1; j <= t2Len; j++ {
			if text1[i-1] == text2[j-1] {
				dp[k][j] = dp[l][j-1] + 1
			} else {
				dp[k][j] = MaxInt(dp[l][j], dp[k][j-1])
			}
		}
		l,k = k,l
	}

	return dp[l][t2Len]
}

func MaxInt(vars ...int) int {
	max := vars[0]

	for _, i := range vars {
		if max < i {
			max = i
		}
	}

	return max
}
```

### [72\. 编辑距离](https://leetcode-cn.com/problems/edit-distance/)

Difficulty: **困难**

给你两个单词 `word1` 和 `word2`，请你计算出将 `word1` 转换成 `word2`所使用的最少操作数 。

你可以对一个单词进行如下三种操作：

* 插入一个字符
* 删除一个字符
* 替换一个字符

**示例 1：**

```
输入：word1 = "horse", word2 = "ros"
输出：3
解释：
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')
```

**示例 2：**

```
输入：word1 = "intention", word2 = "execution"
输出：5
解释：
intention -> inention (删除 't')
inention -> enention (将 'i' 替换为 'e')
enention -> exention (将 'n' 替换为 'x')
exention -> exection (将 'n' 替换为 'c')
exection -> execution (插入 'u')
```

**提示：**

* `0 <= word1.length, word2.length <= 500`
* `word1` 和 `word2` 由小写英文字母组成

#### Solution

Language: ****

```
​func minDistance(word1 string, word2 string) int {
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
	k,l := 0,1
	for i := 1; i <= t1Len; i++ {
		dp[k][0] = i
		for j := 1; j <= t2Len; j++ {
			if word1[i-1] == word2[j-1] {
				dp[k][j] = dp[l][j-1]
			} else {
				dp[k][j] = MinInt(dp[l][j], dp[k][j-1], dp[l][j-1]) + 1
			}
		}
		l,k = k,l
	}

	return dp[l][t2Len]
}

func MinInt(vars ...int) int {
	min := vars[0]

	for _, i := range vars {
		if min > i {
			min = i
		}
	}

	return min
}

```

### [322\. 零钱兑换](https://leetcode-cn.com/problems/coin-change/)

Difficulty: **中等**

给定不同面额的硬币 `coins` 和一个总金额 `amount`。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 `-1`。

你可以认为每种硬币的数量是无限的。

**示例 1：**

```
输入：coins = [1, 2, 5], amount = 11
输出：3 
解释：11 = 5 + 5 + 1
```

**示例 2：**

```
输入：coins = [2], amount = 3
输出：-1
```

**示例 3：**

```
输入：coins = [1], amount = 0
输出：0
```

**示例 4：**

```
输入：coins = [1], amount = 1
输出：1
```

**示例 5：**

```
输入：coins = [1], amount = 2
输出：2
```

**提示：**

* `1 <= coins.length <= 12`
* `1 <= coins[i] <= 2<sup>31</sup> - 1`
* `0 <= amount <= 10<sup>4</sup>`

#### Solution

Language: ****

```
​func coinChange(coins []int, amount int) int {
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
			}else {
				if *min > a+1 {
					*min = a+1
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

func MinInt(vars ...int) int {
	min := vars[0]

	for _, i := range vars {
		if min > i {
			min = i
		}
	}

	return min
}
```

### [120\. 三角形最小路径和](https://leetcode-cn.com/problems/triangle/)

Difficulty: **中等**


给定一个三角形 `triangle` ，找出自顶向下的最小路径和。

每一步只能移动到下一行中相邻的结点上。**相邻的结点** 在这里指的是 **下标** 与 **上一层结点下标** 相同或者等于 **上一层结点下标 + 1** 的两个结点。也就是说，如果正位于当前行的下标 `i` ，那么下一步可以移动到下一行的下标 `i` 或 `i + 1` 。

**示例 1：**

```
输入：triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
输出：11
解释：如下面简图所示：
   2
  3 4
 6 5 7
4 1 8 3
自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
```

**示例 2：**

```
输入：triangle = [[-10]]
输出：-10
```

**提示：**

*   `1 <= triangle.length <= 200`
*   `triangle[0].length == 1`
*   `triangle[i].length == triangle[i - 1].length + 1`
*   `-10<sup>4</sup> <= triangle[i][j] <= 10<sup>4</sup>`

**进阶：**

*   你可以只使用 `O(n)` 的额外空间（`n` 为三角形的总行数）来解决这个问题吗？


#### Solution

Language: ****

```
​func minimumTotal(triangle [][]int) int {
    n := len(triangle)
    f := make([]int, n)
    f[0] = triangle[0][0]
    for i := 1; i < n; i++ {
        f[i] = f[i - 1] + triangle[i][i]
        for j := i - 1; j > 0; j-- {
            f[j] = min(f[j - 1], f[j]) + triangle[i][j]
        }
        f[0] += triangle[i][0]
    }
    ans := math.MaxInt32
    for i := 0; i < n; i++ {
        ans = min(ans, f[i])
    }
    return ans
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

```

### [104\. 二叉树的最大深度](https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/)

Difficulty: **简单**


给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

**说明:** 叶子节点是指没有子节点的节点。

**示例：**  
给定二叉树 `[3,9,20,null,null,15,7]`，

```
    3
   / \
  9  20
    /  \
   15   7
```

返回它的最大深度 3 。


#### Solution

Language: ****

```
​/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return MaxInt(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func MaxInt(vars ...int) int {
	max := vars[0]

	for _, i := range vars {
		if max < i {
			max = i
		}
	}

	return max
}
```

### [300\. 最长递增子序列](https://leetcode-cn.com/problems/longest-increasing-subsequence/)

Difficulty: **中等**


给你一个整数数组 `nums` ，找到其中最长严格递增子序列的长度。

子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，`[3,6,2,7]` 是数组 `[0,3,1,6,2,2,7]` 的子序列。

**示例 1：**

```
输入：nums = [10,9,2,5,3,7,101,18]
输出：4
解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
```

**示例 2：**

```
输入：nums = [0,1,0,3,2,3]
输出：4
```

**示例 3：**

```
输入：nums = [7,7,7,7,7,7,7]
输出：1
```

**提示：**

*   `1 <= nums.length <= 2500`
*   `-10<sup>4</sup> <= nums[i] <= 10<sup>4</sup>`

**进阶：**

*   你可以设计时间复杂度为 `O(n<sup>2</sup>)` 的解决方案吗？
*   你能将算法的时间复杂度降低到 `O(n log(n))` 吗?


#### Solution

Language: ****

```
​func lengthOfLIS(nums []int) int {
	dp := []int{^(1 << 32)}
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
```

### [110\. 平衡二叉树](https://leetcode-cn.com/problems/balanced-binary-tree/)

Difficulty: **简单**


给定一个二叉树，判断它是否是高度平衡的二叉树。

本题中，一棵高度平衡二叉树定义为：

> 一个二叉树_每个节点 _的左右两个子树的高度差的绝对值不超过 1 。

**示例 1：**

![](https://assets.leetcode.com/uploads/2020/10/06/balance_1.jpg)

```
输入：root = [3,9,20,null,null,15,7]
输出：true
```

**示例 2：**

![](https://assets.leetcode.com/uploads/2020/10/06/balance_2.jpg)

```
输入：root = [1,2,2,3,3,null,null,4,4]
输出：false
```

**示例 3：**

```
输入：root = []
输出：true
```

**提示：**

*   树中的节点数在范围 `[0, 5000]` 内
*   `-10<sup>4</sup> <= Node.val <= 10<sup>4</sup>`


#### Solution

Language: ****

```
​/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	b, _ := deepIsBalanced(root)
	return b
}

func deepIsBalanced(root *TreeNode) (bool, int) {
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
	return left-right <= 1 && left-right >= -1, MaxInt(left, right) + 1
}

func MaxInt(vars ...int) int {
	max := vars[0]

	for _, i := range vars {
		if max < i {
			max = i
		}
	}

	return max
}
```

### [124\. 二叉树中的最大路径和](https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/)

Difficulty: **困难**


**路径** 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。同一个节点在一条路径序列中 **至多出现一次** 。该路径 **至少包含一个** 节点，且不一定经过根节点。

**路径和** 是路径中各节点值的总和。

给你一个二叉树的根节点 `root` ，返回其 **最大路径和** 。

**示例 1：**

![](https://assets.leetcode.com/uploads/2020/10/13/exx1.jpg)

```
输入：root = [1,2,3]
输出：6
解释：最优路径是 2 -> 1 -> 3 ，路径和为 2 + 1 + 3 = 6
```

**示例 2：**

![](https://assets.leetcode.com/uploads/2020/10/13/exx2.jpg)

```
输入：root = [-10,9,20,null,null,15,7]
输出：42
解释：最优路径是 15 -> 20 -> 7 ，路径和为 15 + 20 + 7 = 42
```

**提示：**

*   树中节点数目范围是 `[1, 3 * 10<sup>4</sup>]`
*   `-1000 <= Node.val <= 1000`


#### Solution

Language: ****

```
​
```