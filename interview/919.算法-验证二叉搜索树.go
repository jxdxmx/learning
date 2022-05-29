package interview

import "math"

//https://leetcode.cn/problems/validate-binary-search-tree
//98. 验证二叉搜索树

//给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
//有效 二叉搜索树定义如下：
//节点的左子树只包含 小于 当前节点的数。
//节点的右子树只包含 大于 当前节点的数。
//所有左子树和右子树自身必须也是二叉搜索树。
//输入：root = [2,1,3]
//输出：true

type TreeNode919 struct {
	Val   int
	Left  *TreeNode919
	Right *TreeNode919
}

/**
 * Definition for a binary tree node.
 * type TreeNode919 struct {
 *     Val int
 *     Left *TreeNode919
 *     Right *TreeNode919
 * }
 */
func isValidBST919(root *TreeNode919) bool {
	_, _, ans := dfs919(root)
	return ans
}

func dfs919(root *TreeNode919) (minVal, maxVal int64, valid bool) {
	if root == nil {
		return math.MaxInt64, math.MinInt64, true
	}
	minVal1, maxVal1, v1 := dfs919(root.Left)
	minVal2, maxVal2, v2 := dfs919(root.Right)
	v := int64(root.Val)
	if v > maxVal1 && v < minVal2 && v1 && v2 {
		return min919(v, minVal1, minVal2), max919(maxVal2, maxVal1, v), true
	}
	return 0, 0, false
}

func min919(arr ...int64) int64 {
	var ans int64 = math.MaxInt64
	for _, v := range arr {
		if ans > v {
			ans = v
		}
	}
	return ans
}

func max919(arr ...int64) int64 {
	var ans int64 = math.MinInt64
	for _, v := range arr {
		if ans < v {
			ans = v
		}
	}
	return ans
}
