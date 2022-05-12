package interview

func F024() {

}

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

//思路：hash算法、KMP模式匹配算法
