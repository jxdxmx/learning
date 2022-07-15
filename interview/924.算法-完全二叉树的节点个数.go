package interview

// 2022-07-15

//https://leetcode.cn/problems/count-complete-tree-nodes/
//222. 完全二叉树的节点个数
//给你一棵 完全二叉树 的根节点 root ，求出该树的节点个数。
//完全二叉树 的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，
//并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第 h 层，则该层包含 1~ 2h 个节点。

type TreeNode924 struct {
	Val   int
	Left  *TreeNode924
	Right *TreeNode924
}

//示例 1：
//输入：root = [1,2,3,4,5,6]
//输出：6
//
//示例 2：
//输入：root = []
//输出：0
//
//示例 3：
//输入：root = [1]
//输出：1
/**
 * Definition for a binary tree node.
 * type TreeNode925 struct {
 *     Val int
 *     Left *TreeNode925
 *     Right *TreeNode925
 * }
 */

func countNodes924(root *TreeNode924) int {
	if root == nil {
		return 0
	}
	leftN, rightN := countNodes924(root.Left), countNodes924(root.Right)
	return leftN + rightN + 1
}
