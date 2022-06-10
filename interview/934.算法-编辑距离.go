package interview

import "math"

//https://leetcode.cn/problems/edit-distance/
//72. 编辑距离
//给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。
//你可以对一个单词进行如下三种操作：
//插入一个字符
//删除一个字符
//替换一个字符

//示例 1：
//输入：word1 = "horse", word2 = "ros"
//输出：3
//解释：
//horse -> rorse (将 'h' 替换为 'r')
//rorse -> rose (删除 'r')
//rose -> ros (删除 'e')

//示例 2：
//输入：word1 = "intention", word2 = "execution"
//输出：5
//解释：
//intention -> inention (删除 't')
//inention -> enention (将 'i' 替换为 'e')
//enention -> exention (将 'n' 替换为 'x')
//exention -> exection (将 'n' 替换为 'c')
//exection -> execution (插入 'u')

func minDistance934(word1 string, word2 string) int {
	word1 = " " + word1
	word2 = " " + word2
	m, n := len(word1), len(word2)
	var f [][]int
	for i := 0; i < m; i++ {
		f = append(f, make([]int, n))
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				f[0][j] = j
			} else if j == 0 {
				f[i][0] = i
			} else {
				f[i][j] = math.MaxInt32
			}
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if word1[i] == word2[j] {
				f[i][j] = f[i-1][j-1]
			} else {
				// f[i][j] = min934(f[i-1][j]+1(删除字符),f[i][j-1]+1(删除字符),f[i-1][j-1](当前字符相同),f[i-1][j-1]+1(替换当前字符))
				f[i][j] = min934(f[i][j], f[i-1][j]+1, f[i][j-1]+1, f[i-1][j-1]+1)
			}
		}
	}
	return f[m-1][n-1]
}

func min934(arr ...int) int {
	var ans = math.MaxInt32
	for _, n := range arr {
		if n < ans {
			ans = n
		}
	}
	return ans
}
