package interview

import "math/rand"

// 2022-07-11

//https://leetcode.cn/problems/kth-largest-element-in-an-array/
//215. 数组中的第K个最大元素
//
//给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
//请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
//
//示例 1:
//输入: [3,2,1,5,6,4] 和 k = 2
//输出: 5
//
//示例2:
//输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
//输出: 4

func findKthLargest215(nums []int, k int) int {
	ans = 0
	quickSort215(nums, 0, len(nums)-1, k)
	// fmt.Println(nums)
	// return -1
	return ans
}

var ans int

func quickSort215(nums []int, left, right, k int) {
	if left == right && len(nums)-left == k {
		ans = nums[right]
		return
	} else if left > right {
		return
	}
	pivot := partition215(nums, left, right)
	// fmt.Println(nums,"left/right/pivot:",left,right,pivot,"pivotVal:",nums[pivot],nums[left:right+1])
	if len(nums)-pivot == k {
		ans = nums[pivot]
	} else if len(nums)-pivot < k {
		quickSort215(nums, left, pivot, k)
	} else {
		quickSort215(nums, pivot+1, right, k)
	}
}

// [1,1,1,1,1,1,1,2,3,4,3,2,2,2,5,6,45,6,43,43,6,5,6,51,2,3,4,3,2,2,2,5,6,45,6,43,43,6,5,6,5]
// 2
// 为什么for left<=right,而不是< ?
// <时
// [3 2 3 2 2]
// [2 2 3 2 3] pivot/pivotVal: 15 2 i/j 12 16 left/right: 14 14
// <=时，结果正常，返回right 13
// [3 2 3 2 2]
// [2 2 3 2 3] pivot/pivotVal: 15 2 i/j 12 16 left/right: 14 13
// i,j := left,right
// fmt.Println(nums[i:j+1])
// fmt.Println(nums[i:j+1],"pivot/pivotVal:",pivot,pivotVal,"i/j",i,j,"left/right:",left,right)

// 为什么只能left~pivot,而不是left~pivot-1?
// quickSort215(nums,left,pivot)
// [1 1 1 1 1 1 1 2 3 4 3 2 2 2 5 5 6 5 6 6 6 5 5 2 2 3 4 3 2 2 51 6 6 45 43 43 43 43 6 45 6]
//  left/right/pivot: 0 40 29 pivotVal: 2 [1 1 1 1 1 1 1 2 3 4 3 2 2 2 5 5 6 5 6 6 6 5 5 2 2 3 4 3 2 2 51 6 6 45 43 43 43 43 6 45 6]
// 第一次排序，是以51为pivot，返回的index却是51前面的2（29）的index。
// 如果left~pivot-1，就错了,2就漏掉了。应该是left~viot（包含2）
// 但是要注意，也可能pivot也可能是51的index。
// 此时left~pivot不仅包含了51,还包含了51.这是没问题的。

func partition215(nums []int, left, right int) int {
	pivot := rand.Intn(right-left+1) + left
	pivotVal := nums[pivot]
	for left <= right {
		for nums[left] < pivotVal {
			left++
		}
		for nums[right] > pivotVal {
			right--
		}
		if left <= right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}
	return right
}
