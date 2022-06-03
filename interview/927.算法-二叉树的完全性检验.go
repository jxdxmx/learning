package interview

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
	dept := getDepth927(root)
	// fmt.Println(dept)
	var queue []*TreeNode927
	var dm = make(map[*TreeNode927]int)
	dm[root] = 1
	queue = append(queue, root)
	flag := false
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		if dm[top] < dept-1 {
			if top.Left == nil || top.Right == nil {
				return false
			}
			queue = append(queue, top.Left)
			queue = append(queue, top.Right)
			dm[top.Left] = dm[top] + 1
			dm[top.Right] = dm[top] + 1
		} else if !flag && dm[top] == dept-1 {
			if top.Right != nil && top.Left == nil {
				return false
			}
			if top.Right == nil {
				flag = true
			}
		} else if flag && dm[top] == dept-1 {
			if top.Right != nil || top.Left != nil {
				return false
			}
		}
	}
	return true
}

func getDepth927(root *TreeNode927) int {
	if root == nil {
		return 0
	}
	ld := getDepth927(root.Left)
	rd := getDepth927(root.Right)
	if ld < rd {
		ld = rd
	}
	return ld + 1
}
