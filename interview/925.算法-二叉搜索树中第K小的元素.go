package interview

// 2022-07-15

//https://leetcode.cn/problems/kth-smallest-element-in-a-bst/
//230. 二叉搜索树中第K小的元素
//给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。
//示例 1：
//输入：root = [3,1,4,null,2], k = 1
//输出：1
//
//示例 2：
//输入：root = [5,3,6,2,4,null,null,1], k = 3
//输出：3
//

type TreeNode925 struct {
	Val   int
	Left  *TreeNode925
	Right *TreeNode925
}

/**
 * Definition for a binary tree node.
 * type TreeNode925 struct {
 *     Val int
 *     Left *TreeNode925
 *     Right *TreeNode925
 * }
 */
func kthSmallest925(root *TreeNode925, k int) int {
	ans, _ := dfs925(root, k)
	if ans == -1 {
		return 0
	}
	return ans
}

func dfs925(root *TreeNode925, k int) (int, int) { // 要注意，k必须作为参数！！！
	if root == nil {
		return -1, k
	}
	var l, r = -1, -1
	if root.Left != nil {
		l, k = dfs925(root.Left, k)
		if l > 0 {
			return l, k
		}
	}

	k--
	// fmt.Println("k:",k," - ",root.Val)
	if k == 0 {
		return root.Val, k
	}

	if root.Right != nil {
		r, k = dfs925(root.Right, k)
		if r > 0 {
			return r, k
		}
	}
	return -1, k
}
