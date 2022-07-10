package interview

// 2022-07-10

//https://leetcode.cn/problems/find-the-duplicate-number/
//287. 寻找重复数

//给定一个包含n + 1 个整数的数组nums ，其数字都在[1, n]范围内（包括 1 和 n），可知至少存在一个重复的整数。
//假设 nums 只有 一个重复的整数 ，返回这个重复的数 。
//你设计的解决方案必须 不修改 数组 nums 且只用常量级 O(1) 的额外空间。

//示例 1：
//输入：nums = [1,3,4,2,2]
//输出：2

//示例 2：
//输入：nums = [3,1,3,4,2]
//输出：3

// 二分
// 如果可以替换，那么可以修改原数组。nums[i]放入nums[i]位置上。如果nums[i] == nums[nums[i]]，表示这个是重复数字

// 二分
func findDuplicate287(nums []int) int {
	n := len(nums) - 1
	left, right := 1, n
	for left < right {
		mid := (left + right + 1) / 2
		cnt := 0
		for i := 0; i <= n; i++ {
			if nums[i] < mid {
				cnt++
			}
		}
		if cnt < mid {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return right
}

// 如果可以替换，那么可以修改原数组。nums[i]放入nums[i]位置上。如果nums[i] == nums[nums[i]]，表示这个是重复数字
func findDuplicate(nums []int) int {
	//times := 0
	for i := 1; i <= len(nums); i++ {
		// times++
		// fmt.Println(i,1111,nums)
		if nums[i-1] != i {
			if nums[i-1] == nums[nums[i-1]-1] {
				// fmt.Println(len(nums),times)
				return nums[i-1]
			} else {
				nums[i-1], nums[nums[i-1]-1] = nums[nums[i-1]-1], nums[i-1]
			}
			i-- // 这里，交互好以后，需要继续对i做判断，所以i--,for会进行i++
		}
		// fmt.Println(i,2222,nums)
	}
	return -1
}
