package interview

//https://leetcode.cn/problems/n-queens
//51. N 皇后

//按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。
//n皇后问题 研究的是如何将 n个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
//给你一个整数 n ，返回所有不同的n皇后问题 的解决方案。
//每一种解法包含一个不同的n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
//
//示例 1：
//输入：n = 4
//输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
//解释：如上图所示，4 皇后问题存在两个不同的解法。
//
//示例 2：
//输入：n = 1
//输出：[["Q"]]
func solveNQueens922(n int) [][]string {
	var used = make([]bool, n)
	var m1 = make(map[int]bool)
	var m2 = make(map[int]bool)
	var ans [][]int
	var s []int
	var dfs func(row int)
	dfs = func(row int) {
		if row >= n {
			var tmp = make([]int, n)
			copy(tmp, s)
			ans = append(ans, tmp)
			return
		}
		for col := 0; col < n; col++ {
			if used[col] || m1[row+col] || m2[row-col] {
				continue
			}
			used[col] = true
			m1[row+col] = true
			m2[row-col] = true
			s = append(s, col)
			dfs(row + 1)
			used[col] = false
			m1[row+col] = false
			m2[row-col] = false
			s = s[:len(s)-1]
		}
	}
	dfs(0)
	// fmt.Println(ans)
	var result = make([][]string, len(ans))
	for i := 0; i < len(ans); i++ {
		for j := 0; j < n; j++ {
			var bs []byte
			for k := 0; k < n; k++ {
				bs = append(bs, '.')
			}
			bs[ans[i][j]] = 'Q'
			result[i] = append(result[i], string(bs))
		}
	}
	return result
}
