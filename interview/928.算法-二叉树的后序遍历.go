package interview

/**
 * Definition for a binary tree node.
 * type TreeNode928 struct {
 *     Val int
 *     Left *TreeNode928
 *     Right *TreeNode928
 * }
 */

type TreeNode928 struct {
	Val   int
	Left  *TreeNode928
	Right *TreeNode928
}

func postorderTraversal928(root *TreeNode928) []int {
	ans928 = []int{}
	dfs(root)
	return ans928
}

var ans928 []int

func dfs(root *TreeNode928) {
	if root == nil {
		return
	}
	dfs(root.Left)
	dfs(root.Right)
	ans928 = append(ans928, root.Val)
}

// 迭代方式求解
//https://zhuanlan.zhihu.com/p/80578741 二叉树后序遍历的两种易写的非递归写法 - 知乎

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal928_2(root *TreeNode928) []int {
	var res []int
	var stack []*TreeNode928
	stack = append(stack, root)
	var curr *TreeNode928 = root
	for len(stack) > 0 {
		for curr != nil {
			res = append(res, curr.Val)
			stack = append(stack, curr)
			curr = curr.Right
		}
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if top != nil && top.Left != nil {
			curr = top.Left
		}
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return res
}
