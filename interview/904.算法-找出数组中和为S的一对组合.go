package interview

//https://leetcode.cn/problems/two-sum/submissions/
//1. 两数之和

//找出数组中和为S的一对组合，找出一组就行
//
//方法一：map
//方法二：排序+双指针
func twoSum1(nums []int, target int) []int {
	var m = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if x, ok := m[target-nums[i]]; ok {
			return []int{x, i}
		}
		m[nums[i]] = i
	}
	return []int{}
}
