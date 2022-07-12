package interview

// 2022-07-12

// https://leetcode.cn/problems/maximum-depth-of-binary-tree
// 104. 二叉树的最大深度
//给定一个二叉树，找出其最大深度。

//二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

//说明:叶子节点是指没有子节点的节点。
//示例：
//给定二叉树 [3,9,20,null,null,15,7]，
//3
/// \
//9  20
///  \
//15   7
//返回它的最大深度3 。

// 自下而上

type TreeNode920 struct {
	Val   int
	Left  *TreeNode920
	Right *TreeNode920
}

func maxDepth920(root *TreeNode920) int {
	if root == nil {
		return 0
	}
	l, r := maxDepth920(root.Left), maxDepth920(root.Right)
	if l < r {
		l = r
	}
	return l + 1
}

// 自上而下

func maxDepth9202(root *TreeNode920) int {
	return depth9202(root, 0)
}
func depth9202(root *TreeNode920, dep int) int {
	if root == nil {
		return dep
	}
	dep++
	l, r := depth9202(root.Left, dep), depth9202(root.Right, dep)
	if l < r {
		return r
	}
	return l
}
