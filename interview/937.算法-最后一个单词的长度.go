package interview

// 2022-07-17

//https://leetcode.cn/problems/length-of-last-word/
//58. 最后一个单词的长度
//给你一个字符串 s，由若干单词组成，单词前后用一些空格字符隔开。返回字符串中 最后一个 单词的长度。
//单词 是指仅由字母组成、不包含任何空格字符的最大子字符串。
//
//示例 1：
//输入：s = "Hello World"
//输出：5
//解释：最后一个单词是“World”，长度为5。
//
//示例 2：
//输入：s = "   fly me   to   the moon  "
//输出：4
//解释：最后一个单词是“moon”，长度为4。
//
//示例 3：
//输入：s = "luffy is still joyboy"
//输出：6
//解释：最后一个单词是长度为6的“joyboy”。
func lengthOfLastWord937(s string) int {
	var start, end = 0, 0
	bs := []byte(s)
	for i := len(bs) - 1; i >= 0; i-- {
		if (bs[i] >= 'a' && bs[i] <= 'z') || (bs[i] >= 'A' && bs[i] <= 'Z') {
			end = i
			break
		}
	}
	for i := end; i >= 0; i-- {
		if (bs[i] >= 'a' && bs[i] <= 'z') || (bs[i] >= 'A' && bs[i] <= 'Z') {
			start = i
		} else {
			break
		}
	}
	return end - start + 1
}
