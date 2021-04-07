# 剑指 Offer

## 目录

---

<!--ts-->
   * [剑指 Offer](#剑指-offer)
      * [目录](#目录)
         * [<a href="https://leetcode-cn.com/problems/ju-zhen-zhong-de-lu-jing-lcof/" rel="nofollow">剑指 Offer 12. 矩阵中的路径</a>](#剑指-offer-12-矩阵中的路径)
            * [Solution](#solution)
         * [<a href="https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/" rel="nofollow">剑指 Offer 13. 机器人的运动范围</a>](#剑指-offer-13-机器人的运动范围)
            * [Solution](#solution-1)

<!-- Added by: runner, at: Wed Apr  7 14:47:17 UTC 2021 -->

<!--te-->

---

### [剑指 Offer 12\. 矩阵中的路径](https://leetcode-cn.com/problems/ju-zhen-zhong-de-lu-jing-lcof/)

Difficulty: **中等**

请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有字符的路径。路径可以从矩阵中的任意一格开始，每一步可以在矩阵中向左、右、上、下移动一格。如果一条路径经过了矩阵的某一格，那么该路径不能再次进入该格子。例如，在下面的 3×4
的矩阵中包含一条字符串“bfce”的路径（路径中的字母用加粗标出）。

[["a","**b**","c","e"],  
["s","**f**","**c**","s"],  
["a","d","**e**","e"]]

但矩阵中不包含字符串“abfb”的路径，因为字符串的第一个字符 b 占据了矩阵中的第一行第二个格子之后，路径不能再次进入这个格子。

**示例 1：**

```
输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
输出：true
```

**示例 2：**

```
输入：board = [["a","b"],["c","d"]], word = "abcd"
输出：false
```

**提示：**

- `1 <= board.length <= 200`
- `1 <= board[i].length <= 200`

注意：本题与主站 79 题相同：

#### Solution

Language: **GO**

```go
package main

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
	var base = make([][]byte, len(board))
	for y := 0; y < lengthY; y++ {
		for x := 0; x < lengthX; x++ {
			if w[0] == board[y][x] {
				copy(base, board)
				for i := range base {
					base[i] = make([]byte, len(board[i]))
					copy(base[i], board[i])
				}
				if searchString(base, w[1:], x, y) {
					return true
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
	var base = make([][]byte, len(board))
	board[y][x] = ' '
	if x != 0 {
		if board[y][x-1] == word[0] {
			copy(base, board)
			for i := range base {
				base[i] = make([]byte, len(board[i]))
				copy(base[i], board[i])
			}
			if searchString(base, word[1:], x-1, y) {
				return true
			}
		}
	}

	if x != lengthX-1 {
		if board[y][x+1] == word[0] {
			copy(base, board)
			for i := range base {
				base[i] = make([]byte, len(board[i]))
				copy(base[i], board[i])
			}
			if searchString(base, word[1:], x+1, y) {
				return true
			}
		}
	}

	if y != 0 {
		if board[y-1][x] == word[0] {
			copy(base, board)
			for i := range base {
				base[i] = make([]byte, len(board[i]))
				copy(base[i], board[i])
			}
			if searchString(base, word[1:], x, y-1) {
				return true
			}
		}
	}

	if y != lengthY-1 {
		if board[y+1][x] == word[0] {
			copy(base, board)
			for i := range base {
				base[i] = make([]byte, len(board[i]))
				copy(base[i], board[i])
			}
			if searchString(base, word[1:], x, y+1) {
				return true
			}
		}
	}

	return false
}
```

### [剑指 Offer 13\. 机器人的运动范围](https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/)

Difficulty: **中等**

地上有一个 m 行 n 列的方格，从坐标 `[0,0]` 到坐标 `[m-1,n-1]` 。一个机器人从坐标 `[0, 0]` 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于
k 的格子。例如，当 k 为 18 时，机器人能够进入方格 [35, 37] ，因为 3+5+3+7=18。但它不能进入方格 [35, 38]，因为 3+5+3+8=19。请问该机器人能够到达多少个格子？

**示例 1：**

```
输入：m = 2, n = 3, k = 1
输出：3
```

**示例 2：**

```
输入：m = 3, n = 1, k = 0
输出：1
```

**提示：**

- `1 <= n,m <= 100`
- `0 <= k <= 20`

#### Solution

Language: **GO**

```go
package main

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
```
