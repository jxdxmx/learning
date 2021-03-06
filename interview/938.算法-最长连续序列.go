package interview

// 2022-07-17

//https://leetcode.cn/problems/longest-consecutive-sequence/
//128. 最长连续序列
//给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
//请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
//
//示例 1：
//输入：nums = [100,4,200,1,3,2]
//输出：4
//解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
//
//示例 2：
//输入：nums = [0,3,7,2,5,8,4,6,0,1]
//输出：9
func longestConsecutive938(nums []int) int {
	n := len(nums)
	fa := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		fa[i] = i
	}
	var m = make(map[int]int)
	var ans int
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			continue
		}
		m[nums[i]] = i
		if idx1, ok := m[nums[i]-1]; ok {
			union938(fa, i, idx1)
		}
		if idx2, ok := m[nums[i]+1]; ok {
			union938(fa, i, idx2)
		}
	}
	for i := 0; i < n; i++ {
		x := find938(fa, i)
		cnt[x]++
		if cnt[x] > ans {
			ans = cnt[x]
		}
	}
	return ans
}
func find938(fa []int, a int) int {
	if fa[a] == a {
		return a
	}
	fa[a] = find938(fa, fa[a])
	return fa[a]
}
func union938(fa []int, a, b int) {
	if fa[a] == fa[b] {
		return
	}
	fa[a] = find938(fa, b)
}
