package interview

//https://leetcode.cn/problems/intersection-of-two-linked-lists/
//160. 相交链表

//给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。
//图示两个链表在节点 c1 开始相交：
//输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,6,1,8,4,5], skipA = 2, skipB = 3
//输出：Intersected at '8'

type ListNode917 struct {
	Val  int
	Next *ListNode917
}

/**
 * Definition for singly-linked list.
 * type ListNode917 struct {
 *     Val int
 *     Next *ListNode917
 * }
 */
func getIntersectionNode917(headA, headB *ListNode917) *ListNode917 {
	var m = make(map[*ListNode917]bool)
	for headA != nil || headB != nil {
		if headA != nil {
			if _, ok := m[headA]; ok {
				return headA
			}
			m[headA] = true
			headA = headA.Next
		}
		if headB != nil {
			if _, ok := m[headB]; ok {
				return headB
			}
			m[headB] = true
			headB = headB.Next
		}
	}
	return nil
}
