package interview

// 2022-07-09

//https://leetcode.cn/problems/reverse-string/
//344. 反转字符串
//问题描述
//请实现一个算法,在不使用"额外数据结构和储存空间"的情况下,翻转一个给定的字符串(可以使用单个过程变量).
//给定一个string,请返回一个string,为翻转后的字符串.保证字符串的长度小于等于5000.
//
//解题思路
//翻转字符串其实是将一个字符串以中间字符为轴,前后翻转,即将str[len]赋值给str[0],
//将str[0] 赋值 str[len].
//源码参考

func reverseString344(s []byte) {
	n := len(s)
	for i := 0; i < n/2; i++ {
		s[i], s[n-1-i] = s[n-1-i], s[i]
	}
}
