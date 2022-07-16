package interview

// 2022-07-16

//https://leetcode.cn/problems/unique-paths-ii/
//63. 不同路径 II
//一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
//机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish”）。
//现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
//网格中的障碍物和空位置分别用 1 和 0 来表示。
//
//示例 1：
//输入：obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
//输出：2
//解释：3x3 网格的正中间有一个障碍物。
//从左上角到右下角一共有 2 条不同的路径：
//1. 向右 -> 向右 -> 向下 -> 向下
//2. 向下 -> 向下 -> 向右 -> 向右
//
//示例 2：
//输入：obstacleGrid = [[0,1],[0,0]]
//输出：1

//注意：[[0]] = 1  [[1]] = 0
func uniquePathsWithObstacles930(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	var f [][]int
	for i := 0; i < n; i++ {
		f = append(f, make([]int, m))
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				f[i][j] = 0
			} else if i == 0 && j == 0 {
				f[i][j] = 1
			} else if i == 0 {
				f[i][j] = f[i][j-1]
			} else if j == 0 {
				f[i][j] = f[i-1][j]
			} else {
				f[i][j] = f[i-1][j] + f[i][j-1]
			}
		}
	}
	return f[n-1][m-1]
}
