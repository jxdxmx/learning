package interview

// 2022-07-15

//https://leetcode.cn/problems/check-completeness-of-a-binary-tree/
//958. 二叉树的完全性检验
//给定一个二叉树的 root ，确定它是否是一个 完全二叉树 。
//在一个 完全二叉树 中，除了最后一个关卡外，所有关卡都是完全被填满的，并且最后一个关卡中的所有节点都是尽可能靠左的。它可以包含 1 到 2h 节点之间的最后一级 h 。

//示例 1：
//输入：root = [1,2,3,4,5,6]
//输出：true
//解释：最后一层前的每一层都是满的（即，结点值为 {1} 和 {2,3} 的两层），且最后一层中的所有结点（{4,5,6}）都尽可能地向左。

//示例 2：
//输入：root = [1,2,3,4,5,null,7]
//输出：false
//解释：值为 7 的结点没有尽可能靠向左侧。

// 另一种解法：hasNoChild !! 如果false，表示未检测到结尾，true表示已到“结尾”。
// 如果未true，那么，队列中的所有结点，不能有子节点！！
// 如果右节点为空，hasNoChild=true。如果左节点不为空，push队列
// 如果右节点不空，左节点空，return false

type TreeNode927 struct {
	Val   int
	Left  *TreeNode927
	Right *TreeNode927
}

/**
 * Definition for a binary tree node.
 * type TreeNode927 struct {
 *     Val int
 *     Left *TreeNode927
 *     Right *TreeNode927
 * }
 */
func isCompleteTree927(root *TreeNode927) bool {
	var q []*TreeNode927
	q = append(q, root)
	var have = true
	for len(q) > 0 {
		top := q[0]
		q = q[1:]
		if top == nil {
			continue
		}
		if top.Left == nil && top.Right != nil {
			return false
		}
		if !have && (top.Left != nil || top.Right != nil) {
			return false
		}
		if top.Left == nil || (top.Left != nil && top.Right == nil) {
			have = false
		}
		q = append(q, top.Left, top.Right)
	}
	return true
}
