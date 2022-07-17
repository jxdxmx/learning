package interview

// 2022-07-17

import "math"

//https://leetcode.cn/problems/longest-common-prefix/
//14. 最长公共前缀
//编写一个函数来查找字符串数组中的最长公共前缀。
//如果不存在公共前缀，返回空字符串 ""。
//
//示例 1：
//输入：strs = ["flower","flow","flight"]
//输出："fl"
//
//示例 2：
//输入：strs = ["dog","racecar","car"]
//输出：""
//解释：输入不存在公共前缀。

func longestCommonPrefix941(strs []string) string {
	n := math.MaxInt32
	for _, str := range strs {
		if len(str) < n {
			n = len(str)
		}
	}
	var ans = -1
	for i := 0; i < n; i++ {
		flg := false
		b := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if b != strs[j][i] {
				flg = true
				break
			}
		}
		if flg {
			break
		}
		ans = i
	}
	return strs[0][:ans+1]
}
