package interview

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
	_, v := dfs926(root)
	return v
}

func dfs926(root *TreeNode926) (int, bool) {
	if root == nil {
		return 0, true
	}
	d1, v1 := dfs926(root.Left)
	if !v1 {
		return 0, false
	}
	d2, v2 := dfs926(root.Right)
	if !v2 {
		return 0, false
	}
	if d1-d2 > 1 || d2-d1 > 1 {
		return 0, false
	}
	return max(d1, d2) + 1, true
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
