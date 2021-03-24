[TOC]

# Leetcode

### [64\. 最小路径和](https://leetcode-cn.com/problems/minimum-path-sum/)

Difficulty: **中等**


给定一个包含非负整数的 `_m_ x _n_` 网格 `grid` ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

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

*   `m == grid.length`
*   `n == grid[i].length`
*   `1 <= m, n <= 200`
*   `0 <= grid[i][j] <= 100`


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

*   `1 <= m, n <= 100`
*   题目数据保证答案小于等于 `2 * 10<sup>9</sup>`


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

*   `m == obstacleGrid.length`
*   `n == obstacleGrid[i].length`
*   `1 <= m, n <= 100`
*   `obstacleGrid[i][j]` 为 `0` 或 `1`


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