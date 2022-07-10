package interview

import "math"

// 2022-07-10

//https://leetcode.cn/problems/implement-strstr/
//28. 实现 strStr()
//当 needle 是空字符串时，我们应当返回什么值呢？这是⼀个在⾯试中很好的问题。

//题⽬：实现 strStr()
//实现 strStr() 函数。给定⼀个 haystack 字符串和⼀个 needle 字符串，在 haystack 字
//符串中找出 needle 字符串出现的第⼀个位置 (从0开始)。如果不存在，则返回 -1。
//示例 1:
//输⼊: haystack = "hello", needle = "ll" 输出: 2
//示例 2:
//输⼊: haystack = "aaaaa", needle = "bba" 输出: -1
//说明:
//当 needle 是空字符串时，我们应当返回什么值呢？这是⼀个在⾯试中很好的问题。
//对于本题⽽⾔，当 needle 是空字符串时我们应当返回 0 。这与C语⾔的 strstr() 以及
//Java的 indexOf() 定义相符。

//fmt.Println(strings.Index("abc", "")) // 0

//思路：朴素做法、hash算法、KMP模式匹配算法、Sunday匹配

// Sunday匹配算法
// 对⻬⽬标串和模式串，从前向后匹配
// 关注主串中位于模式串后⾯的第⼀个元素（核⼼）
// 如果关注的字符没有在⼦串中出现则直接跳过
// 否则开始移动模式串，移动位数 = ⼦串⻓度 - 该字符最右出现的位置(以0开始)

// hash
func strStr28_1(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	haystack = " " + haystack
	var multi, p, pt uint = 131, 1, 0
	modP := uint(math.MaxInt32 - 1)
	var ps = make([]uint, len(haystack))
	for i := 1; i < len(haystack); i++ {
		ps[i] = (ps[i-1]*multi + uint(haystack[i])) % modP
	}
	for i := 0; i < len(needle); i++ {
		pt = (pt*multi + uint(needle[i])) % modP
		p = (p * multi) % modP
	}
	for i := 1; i <= len(haystack)-len(needle); i++ {
		rb := (ps[i+len(needle)-1] - ps[i-1]*p%modP + modP) % modP
		//fmt.Println(i,i+len(needle),"rb:",rb,"pt:",pt,"p:",p,"modP:",modP)
		if rb == pt {
			return i - 1
		}
	}
	return -1
}

// 朴素做法
func strStr28_2(haystack string, needle string) int {
	m, n := len(haystack), len(needle)
	if n == 0 {
		return 0
	}
	for i := 0; i <= m-n; i++ {
		flag := true
		for j := 0; j < n; j++ {
			if haystack[i+j] != needle[j] {
				flag = false
				break
			}
		}
		if flag {
			return i
		}
	}
	return -1
}

// KMP模式匹配
func strStr28_3(s string, t string) int {
	// KMP模式匹配做法
	var lenS, lenT = len(s), len(t)
	var next = make([]int, lenT)
	// 自匹配
	for i, j := 1, 0; i < lenT; i++ {
		for j > 0 && t[i] != t[j] {
			j = next[j-1]
		}
		if t[i] == t[j] {
			j++
		}
		next[i] = j
	}
	//  fmt.Println(next,lenS)
	for i, j := 0, 0; i < lenS; i++ {
		for j > 0 && s[i] != t[j] {
			j = next[j-1]
		}
		//  fmt.Println("2--","i:",i,"j:",j,"s[i],t[j]:",s[i],t[j])
		if s[i] == t[j] {
			j++
		}
		if j == lenT {
			return i - lenT + 1
		}
	}
	return -1
}

// Sunday匹配算法
// 对⻬⽬标串和模式串，从前向后匹配
// 关注主串中位于模式串后⾯的第⼀个元素（核⼼）
// 如果关注的字符没有在⼦串中出现则直接跳过
// 否则开始移动模式串，移动位数 = ⼦串⻓度 - 该字符最右出现的位置(以0开始)
func strStr028_4(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	var m = make(map[byte]int)
	for i, b := range []byte(needle) {
		m[b] = i
	}
	var n = len(needle)
	for i := 0; i <= len(haystack)-n; { // sunday匹配算法
		j := 0
		for ; j < n; j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		if j == n {
			return i
		}
		// 看i+n
		if i+n < len(haystack) {
			if idx, ok := m[haystack[i+n]]; !ok {
				i += n
			} else {
				i += n - idx
			}
		} else {
			break
		}
	}
	return -1
}
