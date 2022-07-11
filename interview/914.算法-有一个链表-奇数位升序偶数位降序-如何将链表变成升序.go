package interview

// 2022-07-11

import "fmt"

//https://leetcode.cn/problems/merge-two-sorted-lists/
//21. 合并两个有序链表

//有一个链表，奇数位升序偶数位降序，如何将链表变成升序
//思路：
//1、按寄偶位拆分链表
//2、将偶链表反转
//3、合并两个链表

/**
 * Definition for singly-linked list.
 * type ListNode914 struct {
 *     Val int
 *     Next *ListNode914
 * }
 */
type ListNode914 struct {
	Val  int
	Next *ListNode914
}

func mergeTwoLists914(list1 *ListNode914, list2 *ListNode914) *ListNode914 {
	// 按奇偶拆分链表1，并打印它，看对不对
	// printNodeList_914(list1)
	// newList1,newList2 := splitNodeList_914(list1)
	// printNodeList_914(newList1)
	// printNodeList_914(newList2)

	// 反转一个链表（偶），并打印，看对不对
	// printNodeList_914(list1)
	// newList1 := reverseList_914(list1)
	// printNodeList_914(newList1)

	// 合并两个链表-链表1和链表2
	list := mergeTwoLists2_914(list1, list2)

	return list
}

func mergeTwoLists2_914(list1 *ListNode914, list2 *ListNode914) *ListNode914 {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}
	if list1.Val > list2.Val {
		list1, list2 = list2, list1
	}
	head := list1
	for list2 != nil {
		if list1.Next == nil {
			list1.Next = list2
			break
		} else if list1.Next.Val <= list2.Val {
			list1 = list1.Next
		} else {
			tmp := list1.Next
			list1.Next = list2
			list1 = list2
			list2 = tmp
		}
	}
	return head
}

func reverseList_914(head *ListNode914) *ListNode914 {
	var last *ListNode914
	for head != nil {
		tmp := head.Next
		head.Next = last
		last = head
		head = tmp
	}
	return last
}

func print914(head *ListNode914) {
	for head != nil {
		fmt.Print(head.Val, " -> ")
		head = head.Next
	}
	fmt.Println()
}

func split914(head *ListNode914) (*ListNode914, *ListNode914) {
	var l1, l2 *ListNode914 = &ListNode914{}, &ListNode914{}
	p1, p2 := l1, l2
	for head != nil {
		l1.Next = head
		l1 = l1.Next
		head = head.Next
		if head != nil {
			l2.Next = head
			l2 = l2.Next
			head = head.Next
		} else {
			l2.Next = nil
		}
	}
	l1.Next = nil
	return p1.Next, p2.Next
}

func printNodeList_914(root *ListNode914) {
	for root != nil {
		fmt.Print(root.Val)
		if root.Next != nil {
			fmt.Print(",")
		}
		root = root.Next
	}
	fmt.Println()
}
