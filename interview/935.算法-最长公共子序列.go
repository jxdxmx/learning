package interview

//https://leetcode.cn/problems/longest-common-subsequence/
//1143. 最长公共子序列
//给定两个字符串 text1 和 text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。
//一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
//例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
//两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。

//示例 1：
//输入：text1 = "abcde", text2 = "ace"
//输出：3
//解释：最长公共子序列是 "ace" ，它的长度为 3 。
//
//示例 2：
//输入：text1 = "abc", text2 = "abc"
//输出：3
//解释：最长公共子序列是 "abc" ，它的长度为 3 。
//
//示例 3：
//输入：text1 = "abc", text2 = "def"
//输出：0
//解释：两个字符串没有公共子序列，返回 0 。

func longestCommonSubsequence935(text1 string, text2 string) int {
	text1 = " " + text1
	text2 = " " + text2
	m, n := len(text1), len(text2)
	var f [][]int
	for i := 0; i < m; i++ {
		f = append(f, make([]int, n))
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				f[0][j] = 1
			} else if j == 0 {
				f[i][0] = 1
			}
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if text1[i] == text2[j] {
				f[i][j] = f[i-1][j-1] + 1
				// fmt.Println(111,i,j,f[i][j])
			} else {
				// f[i][j] = max935(f[i][j-1],f[i-1][j],f[i-1][j-1](当前不同),f[i-1][j-1]+1(当前相同))
				f[i][j] = max935(f[i][j], f[i-1][j], f[i][j-1], f[i-1][j-1])
				// fmt.Println(i,j,f[i][j])
			}
		}
	}
	return f[m-1][n-1] - 1
}
func max935(arr ...int) int {
	var ans = 0
	for _, n := range arr {
		if n > ans {
			ans = n
		}
	}
	return ans
}
