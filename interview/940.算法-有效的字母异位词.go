package interview

//https://leetcode.cn/problems/valid-anagram/
//242. 有效的字母异位词
//给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
//注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。
//
//示例 1:
//输入: s = "anagram", t = "nagaram"
//输出: true
//
//示例 2:
//输入: s = "rat", t = "car"
//输出: false

func isAnagram940(s string, t string) bool {
	var arr1 [26]int
	var arr2 [26]int
	for _, b := range s {
		arr1[b-'a']++
	}
	for _, b := range t {
		arr2[b-'a']++
	}
	return arr1 == arr2
}
