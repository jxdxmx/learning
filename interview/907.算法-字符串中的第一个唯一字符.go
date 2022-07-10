package interview

// 2022-07-10

//https://leetcode.cn/problems/first-unique-character-in-a-string/

//387. 字符串中的第一个唯一字符
//给定一个字符串,找到 它的第一个不重复的字符，并返回它的索引 。如果不存在，则返回 -1。
//示例 1：
//输入: s = "leetcode"
//输出: 0
//示例 2:
//输入: s = "loveleetcode"
//输出: 2
//示例 3:
//输入: s = "aabb"
//输出: -1

// 这个速度最快！！！
func firstUniqChar387(s string) int {
	var sum [26]int
	for _, b := range []byte(s) {
		sum[b-'a']++
	}
	for i, b := range []byte(s) {
		if sum[b-'a'] == 1 {
			return i
		}
	}
	return -1
}

func firstUniqChar387_3(s string) int {
	var m = make(map[byte]int) // 字母出现次数
	for _, b := range []byte(s) {
		m[b]++
	}
	for i, b := range []byte(s) {
		if m[b] == 1 {
			return i
		}
	}
	return -1
}
func firstUniqChar387_2(s string) int {
	var m = make(map[byte]int)  // 字母出现次数
	var m2 = make(map[byte]int) // 字母最小下标
	for i, b := range []byte(s) {
		m[b]++
		if _, ok := m2[b]; !ok {
			m2[b] = i
		}
	}
	var ans = len(s)
	for k, v := range m {
		if v == 1 {
			if ans > m2[k] {
				ans = m2[k]
			}
		}
	}
	if ans == len(s) {
		return -1
	}
	return ans
}
