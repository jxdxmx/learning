package interview

//https://leetcode.cn/problems/coin-change/
//https://leetcode.cn/problems/coin-change-2/

//322. 零钱兑换
//给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
//计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
//你可以认为每种硬币的数量是无限的。
//
//示例 1：
//输入：coins = [1, 2, 5], amount = 11
//输出：3
//解释：11 = 5 + 5 + 1
//
//示例 2：
//输入：coins = [2], amount = 3
//输出：-1
//
//示例 3：
//输入：coins = [1], amount = 0
//输出：0

func coinChange932(coins []int, amount int) int {
	var f = make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		f[i] = 100000
	}
	for i := 1; i <= amount; i++ {
		for _, c := range coins {
			if i >= c {
				f[i] = min932(f[i], f[i-c]+1)
			}
		}
	}
	if f[amount] >= 100000 {
		return -1
	}
	return f[amount]
}
func min932(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

////518. 零钱兑换 II -- 完全背包问题！
//给你一个整数数组 coins 表示不同面额的硬币，另给一个整数 amount 表示总金额。
//请你计算并返回可以凑成总金额的硬币组合数。如果任何硬币组合都无法凑出总金额，返回 0 。
//假设每一种面额的硬币有无限个。
//题目数据保证结果符合 32 位带符号整数。
//
//示例 1：
//输入：amount = 5, coins = [1, 2, 5]
//输出：4
//解释：有四种方式可以凑成总金额：
//5=5
//5=2+2+1
//5=2+1+1+1
//5=1+1+1+1+1
//
//示例 2：
//输入：amount = 3, coins = [2]
//输出：0
//解释：只用面额 2 的硬币不能凑成总金额 3 。
//
//示例 3：
//输入：amount = 10, coins = [10]
//输出：1
func change932(amount int, coins []int) int {
	var f = make([]int, amount+1)
	f[0] = 1
	for _, c := range coins {
		for i := 1; i <= amount; i++ {
			if i >= c {
				f[i] += f[i-c]
			}
		}
	}
	return f[amount]
}
