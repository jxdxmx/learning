package interview

// 2022-07-16

//https://leetcode.cn/problems/climbing-stairs/
//70. 爬楼梯
//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//
//每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//示例 1：
//输入：n = 2
//输出：2
//解释：有两种方法可以爬到楼顶。
//1. 1 阶 + 1 阶
//2. 2 阶
//
//示例 2：
//输入：n = 3
//输出：3
//解释：有三种方法可以爬到楼顶。
//1. 1 阶 + 1 阶 + 1 阶
//2. 1 阶 + 2 阶
//3. 2 阶 + 1 阶

// 注意：直接使用递归,return climbStairs931(n-1)+climbStairs931(n-2) n=40就会超时
func climbStairs931(n int) int {
	var f = make([]int, n+1)
	f[0] = 1
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}
