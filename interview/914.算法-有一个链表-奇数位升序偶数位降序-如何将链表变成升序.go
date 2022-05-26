package interview

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

func splitNodeList_914(list *ListNode914) (list1 *ListNode914, list2 *ListNode914) {
	if list.Next == nil {
		return list, nil
	}
	list1, list2 = list, list.Next
	head1, head2 := list1, list2
	var next1, next2 *ListNode914
	for {
		if head2.Next == nil {
			next1 = nil
			next2 = nil
			break
		}
		next1 = head2.Next
		if head2.Next.Next == nil {
			next2 = nil
			break
		}
		next2 = next1.Next
		head1.Next = next1
		head2.Next = next2
		head1 = next1
		head2 = next2
	}
	head1.Next = next1
	head2.Next = next2
	return
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
