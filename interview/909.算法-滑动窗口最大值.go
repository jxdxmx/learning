package interview

// 2022-07-10

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

// 思路二：使用最大堆
func maxSlidingWindow909(nums []int, k int) []int {
	var m = make(map[int]*MyNode909)
	var ans []int
	var h Heap909
	for i := 0; i < len(nums); i++ {
		item := m[nums[i]]
		if item == nil {
			item = &MyNode909{val: nums[i], index: i}
			h.Push(item)
		} else {
			item.index = i
		}
		if i >= k-1 {
			for !h.IsEmpty() {
				top := h.top()
				if top.index < i-k+1 {
					h.Pop()
				} else {
					ans = append(ans, top.val)
					break
				}
			}
		}
	}
	return ans
}

type MyNode909 struct {
	val, index int
}
type Heap909 struct {
	items []*MyNode909
}

func (h *Heap909) IsEmpty() bool {
	return len(h.items) == 0
}
func (h *Heap909) Push(item *MyNode909) {
	h.items = append(h.items, item)
	p := len(h.items) - 1
	for p > 0 {
		fa := (p - 1) / 2
		if h.items[fa].val >= h.items[p].val {
			break
		}
		h.items[p], h.items[fa] = h.items[fa], h.items[p]
		p = fa
	}
}
func (h *Heap909) top() *MyNode909 {
	return h.items[0]
}
func (h *Heap909) Pop() *MyNode909 {
	top := h.items[len(h.items)-1]
	h.items[0] = h.items[len(h.items)-1]
	h.items = h.items[:len(h.items)-1]
	p := 0
	ch := 2*p + 1
	for ch < len(h.items) {
		other := 2*p + 2
		if other < len(h.items) && h.items[other].val > h.items[ch].val {
			ch = other
		}
		if h.items[ch].val <= h.items[p].val {
			break
		}
		h.items[p], h.items[ch] = h.items[ch], h.items[p]
		p = ch
		ch = 2*p + 1
	}
	return top
}
