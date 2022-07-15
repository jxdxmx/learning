package interview

// 2022-07-15

//https://leetcode.cn/problems/balanced-binary-tree/
//110. 平衡二叉树
//给定一个二叉树，判断它是否是高度平衡的二叉树。
//本题中，一棵高度平衡二叉树定义为：
//一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。
//示例 1：
//输入：root = [3,9,20,null,null,15,7]
//输出：true
//
//示例 2：
//输入：root = [1,2,2,3,3,null,null,4,4]
//输出：false
//
//示例 3：
//输入：root = []
//输出：true

type TreeNode926 struct {
	Val   int
	Left  *TreeNode926
	Right *TreeNode926
}

/**
 * Definition for a binary tree node.
 * type TreeNode926 struct {
 *     Val int
 *     Left *TreeNode926
 *     Right *TreeNode926
 * }
 */
func isBalanced926(root *TreeNode926) bool {
	_, f := dfs926(root)
	return f
}

func dfs926(root *TreeNode926) (int, bool) {
	if root == nil {
		return 0, true
	}
	l, lf := dfs926(root.Left)
	r, rf := dfs926(root.Right)
	if lf && rf && abs(l-r) <= 1 {
		return max(l, r) + 1, true
	}
	return -1, false
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
