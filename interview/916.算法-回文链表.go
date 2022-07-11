package interview

// 2022-07-11

//https://leetcode.cn/problems/palindrome-linked-list/
//给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。

//示例 1：
//输入：head = [1,2,2,1]
//输出：true

//示例 2：
//输入：head = [1,2]
//输出：false

type ListNode916 struct {
	Val  int
	Next *ListNode916
}

/**
 * Definition for singly-linked list.
 * type ListNode916 struct {
 *     Val int
 *     Next *ListNode916
 * }
 */
func isPalindrome916(head *ListNode916) bool {
	var x = 0
	n, newHead := reverse916(head)
	for x < n/2 {
		if head.Val != newHead.Val {
			return false
		}
		head = head.Next
		newHead = newHead.Next
		x++
	}
	return true
}

func reverse916(head *ListNode916) (int, *ListNode916) {
	var newHead *ListNode916
	var last *ListNode916
	var i int
	for head != nil {
		newHead = &ListNode916{Val: head.Val}
		newHead.Next = last
		last = newHead
		head = head.Next
		i++
	}
	return i, last
}
