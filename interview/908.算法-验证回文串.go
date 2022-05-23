package interview

//125. 验证回文串
//https://leetcode.cn/problems/valid-palindrome/

//给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
//
//说明：本题中，我们将空字符串定义为有效的回文串。
//
//
//
//示例 1:
//
//输入: "A man, a plan, a canal: Panama"
//输出: true
//解释："amanaplanacanalpanama" 是回文串
//示例 2:
//
//输入: "race a car"
//输出: false
//解释："raceacar" 不是回文串
func isPalindrome_125(s string) bool {
	var bs []byte
	for _, b := range []byte(s) {
		if b >= 'a' && b <= 'z' {
			bs = append(bs, b)
		} else if b >= 'A' && b <= 'Z' {
			bs = append(bs, b+'a'-'A')
		} else if b >= '0' && b <= '9' {
			bs = append(bs, b)
		}
	}
	i, j := 0, len(bs)-1
	for i < j {
		if bs[i] != bs[j] {
			return false
		}
		i++
		j--
	}
	return true
}
