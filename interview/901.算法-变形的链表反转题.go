package interview

//https://leetcode.cn/problems/reverse-nodes-in-k-group/
//25. K 个一组翻转链表

//题目
//这其实是一道变形的链表反转题,大致描述如下 给定一个单链表的头节点 head,实现一
//个调整单链表的函数,使得每K个节点之间为一组进行逆序,并且从链表的尾部开始组
//起,头部剩余节点数量不够一组的不需要逆序.(不能使用队列或者栈作为辅助)
//例如:
//链表: 1->2->3->4->5->6->7->8->null, K = 3 .那么 6->7->8 , 3->4->5 , 1->2 各位一组.
//	调整后: 1->2->5->4->3->8->7->6->null .其中 1,2不调整,因为不够一组.
//解析
//原文: https://juejin.im/post/5d4f76325188253b49244dd0

/**
 * Definition for singly-linked list.
 * type ListNode901 struct {
 *     Val int
 *     Next *ListNode901
 * }
 */

type ListNode901 struct {
	Val  int
	Next *ListNode901
}

func reverseKGroup901(head *ListNode901, k int) *ListNode901 {
	var protect *ListNode901
	var last *ListNode901
	tail := getGroup901(head, k)
	if tail == nil {
		return head
	}
	for {
		tail = getGroup901(head, k)
		if tail == nil {
			last.Next = head
			break
		}
		next := tail.Next
		h, t := reverse901(head, tail)
		if last == nil {
			last = t
			protect = h
		} else {
			last.Next = h
			last = t
		}
		head = next
	}
	return protect
}

func getGroup901(head *ListNode901, k int) *ListNode901 {
	k--
	for k > 0 && head != nil {
		head = head.Next
		k--
	}
	if k > 0 {
		return nil
	}
	return head
}

// 反转2
func reverse901(head, tail *ListNode901) (*ListNode901, *ListNode901) {
	h, t := head, tail
	var last *ListNode901
	for head != tail {
		next := head.Next
		head.Next = last
		last = head
		head = next
	}
	head.Next = last
	return t, h
}

///**
// * Definition for singly-linked list.
// * type ListNode025 struct {
// *     Val int
// *     Next *ListNode025
// * }
// */
//// 这个是顺序，1->2->3->4->5变成2->1->4->3->5
//// 如果要逆序，1->2->3->4->5变成1->3->2->5->4,可以先看看总节点数，然后看“多余几个节点”，让head走几步，即可变成"顺序"
//func reverseKGroup901(head *ListNode025, k int) *ListNode025 {
//	var protect,last,end *ListNode025
//	newHeader := getNewHead(head,k)
//	last = &ListNode025{}
//	if newHeader != head {
//		protect = head
//		last = head
//		head = newHeader
//		for last.Next != head {
//			last = last.Next
//		}
//	}
//	fmt.Println("last:",last.Val,"head:",head,"newHeader:",newHeader)
//	end = getGroupEnd025(head,k) // 2
//	if protect == nil {
//		protect = end
//	}
//	for end != nil {
//		next := end.Next // 3
//		newStart,newEnd := reverseTwo025(head,k)
//		last.Next = newStart
//		last = newEnd
//		head = next
//		end = getGroupEnd025(next,k)
//	}
//	if last != head {
//		last.Next = head
//	}
//	return protect
//}
//
//func getNewHead(head *ListNode025,k int) *ListNode025 {
//	var cnt = 1
//	var protect = head
//	for head.Next != nil {
//		cnt++
//		head = head.Next
//	}
//	cnt = cnt % k
//	for cnt > 0 {
//		protect = protect.Next
//		cnt--
//	}
//	return protect
//}
