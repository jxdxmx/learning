package interview

//https://leetcode.cn/problems/longest-palindromic-substring/
//5. 最长回文子串
//给你一个字符串 s，找到 s 中最长的回文子串。

//示例 1：
//输入：s = "babad"
//输出："bab"
//解释："aba" 同样是符合题意的答案。
//
//示例 2：
//输入：s = "cbbd"
//输出："bb"

func longestPalindrome936(s string) string {
	n := len(s)
	var ans = string(s[0])
	var maxLen = 1
	// aba
	for i := 1; i < n-1; i++ {
		l, r := 0, min936(i, n-i-1)
		for l < r {
			mid := (l + r + 1) / 2
			if valid936(s, i-mid, i+mid) {
				l = mid
			} else {
				r = mid - 1
			}
		}
		if r*2+1 > maxLen {
			maxLen = r*2 + 1
			ans = string(s[i-r : i+r+1])
		}
	}
	// bba
	for i := 0; i < n-1; i++ {
		l, r := -1, min936(i, n-i-2)
		for l < r {
			mid := (l + r + 1) / 2
			if valid936(s, i-mid, i+mid+1) {
				l = mid
			} else {
				r = mid - 1
			}
		}
		if r*2+2 > maxLen {
			maxLen = r*2 + 2
			ans = string(s[i-r : i+r+2])
		}
	}
	return ans
}

func min936(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func valid936(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
