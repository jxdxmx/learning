package interview

import "math"

//53. 最大子数组和
//
//https://leetcode.cn/problems/maximum-subarray/

// 思路：前缀和+前缀最小值
func maxSubArray53(nums []int) int {
	var n = len(nums)
	var sum = make([]int, n+1)
	var ans, preMin = math.MinInt32, math.MaxInt32
	for i := 0; i < n; i++ {
		sum[i+1] = sum[i] + nums[i]
	}
	for i := 1; i <= n; i++ {
		preMin = min53(preMin, sum[i-1])
		ans = max53(ans, sum[i]-preMin)
	}
	return ans
}

func max53(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min53(i, j int) int {
	if i > j {
		return j
	}
	return i
}
