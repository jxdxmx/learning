package interview

import "math"

// 2022-07-10

//53. 最大子数组和
//
//https://leetcode.cn/problems/maximum-subarray/

// 思路1：前缀和+前缀最小值
func maxSubArray9051(nums []int) int {
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

// 思路2：动态规划
func maxSubArray9052(nums []int) int {
	var f = make([]int, len(nums))
	f[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if f[i-1] > 0 {
			f[i] = f[i-1] + nums[i]
		} else {
			f[i] = nums[i]
		}
	}
	var ans = f[0]
	for i := 1; i < len(f); i++ {
		if f[i] > ans {
			ans = f[i]
		}
	}
	return ans
}
