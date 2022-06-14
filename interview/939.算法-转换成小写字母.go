package interview

//https://leetcode.cn/problems/to-lower-case/
//709. 转换成小写字母
//给你一个字符串 s ，将该字符串中的大写字母转换成相同的小写字母，返回新的字符串。
//
//示例 1：
//输入：s = "Hello"
//输出："hello"

//示例 2：
//输入：s = "here"
//输出："here"

//示例 3：
//输入：s = "LOVELY"
//输出："lovely"

func toLowerCase939(s string) string {
	bs := []byte(s)
	for i := 0; i < len(bs); i++ {
		if bs[i] >= 'A' && bs[i] <= 'Z' {
			bs[i] += 32
		}
	}
	return string(bs)
}
