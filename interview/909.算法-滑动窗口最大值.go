package interview

//239. 滑动窗口最大值
//https://leetcode.cn/problems/sliding-window-maximum/

//给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
//
//返回 滑动窗口中的最大值 。
//
//
//
//示例 1：
//
//输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
//输出：[3,3,5,5,6,7]
//解释：
//滑动窗口的位置                最大值
//---------------               -----
//[1  3  -1] -3  5  3  6  7       3
//1 [3  -1  -3] 5  3  6  7       3
//1  3 [-1  -3  5] 3  6  7       5
//1  3  -1 [-3  5  3] 6  7       5
//1  3  -1  -3 [5  3  6] 7       6
//1  3  -1  -3  5 [3  6  7]      7
//示例 2：
//
//输入：nums = [1], k = 1
//输出：[1]

// 还可以使用堆+惰性删除
// 直接使用切片--单调队列
func maxSlidingWindow_239_2(nums []int, k int) []int {
	var vals []Item_239
	var ans []int
	n := len(nums)
	for i := 0; i < n; i++ {
		for len(vals) > 0 {
			if vals[0].index < i-k+1 {
				vals = vals[1:]
			} else {
				break
			}
		}
		for len(vals) > 0 {
			if vals[len(vals)-1].val < nums[i] {
				vals = vals[:len(vals)-1]
			} else {
				break
			}
		}
		vals = append(vals, Item_239{val: nums[i], index: i})
		if i >= k-1 {
			ans = append(ans, vals[0].val)
		}
	}

	return ans
}

// 单调队列
func maxSlidingWindow_239(nums []int, k int) []int {
	var h queue_239
	n := len(nums)
	var ans []int
	for i := 0; i < n; i++ {
		var top Item_239
		for !h.IsEmpty() {
			top = h.Top()
			if top.index >= i-k+1 {
				break
			}
			h.PopHead()
		}
		for !h.IsEmpty() {
			tail := h.Tail()
			if tail.val >= nums[i] {
				break
			}
			h.PopTail()
		}
		h.Insert(Item_239{val: nums[i], index: i})
		if i >= k-1 {
			top = h.Top()
			ans = append(ans, top.val)
		}
	}
	return ans
}

type Item_239 struct {
	val   int
	index int
}

type queue_239 struct {
	vals []Item_239
}

func (q *queue_239) IsEmpty() bool {
	return len(q.vals) == 0
}
func (q *queue_239) Top() Item_239 {
	return q.vals[0]
}
func (q *queue_239) Tail() Item_239 {
	return q.vals[len(q.vals)-1]
}
func (q *queue_239) PopHead() {
	q.vals = q.vals[1:]
}
func (q *queue_239) PopTail() {
	q.vals = q.vals[0 : len(q.vals)-1]
}
func (q *queue_239) Insert(item Item_239) {
	q.vals = append(q.vals, item)
}
