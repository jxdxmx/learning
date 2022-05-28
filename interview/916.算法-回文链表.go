package interview

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
	newHead, num := reverse916(head)
	for num/2 > 0 {
		if head.Val != newHead.Val {
			return false
		}
		head = head.Next
		newHead = newHead.Next
		num--
	}
	return true
}

// 反转链表，返回新链表头和链表节点数
func reverse916(head *ListNode916) (*ListNode916, int) {
	if head == nil {
		return nil, 0
	}
	var num = 1
	var last, h2 *ListNode916
	for head != nil {
		h2 = &ListNode916{Val: head.Val}
		h2.Next = last
		last = h2
		num++
		head = head.Next
	}
	return last, num
}
