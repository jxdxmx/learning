package interview

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//多协程查询切片问题
//题目
//假设有一个超长的切片,切片的元素类型为int,切片中的元素为乱序排序.限时5秒,
//使用多个goroutine查找切片中是否存在给定的值,在查找到目标值或者超时后立刻结束所有goroutine的执行.
//比如,切片[23,32,78,43,76,65,345,762,......915,86] ,查找目标值为 345 ,如果切片中存在,
//则目标值输出 "Found it!" 并立即取消仍在执行查询任务的 goroutine.
//如果在超时时间未查到目标值程序,则输出 "Timeout！Not Found" ,同时立即取消仍在执行的查找任务的 goroutine.
//答案: https://mp.weixin.qq.com/s/GhC2WDw3VHP91DrrFVCnag

func F016() {
	var wg sync.WaitGroup
	var slice = []int{23, 32, 78, 43, 76, 65, 345, 762, 915, 86, 23, 32, 78, 43, 76, 65, 345, 762, 915, 86}
	var n = 2
	var target = 333

	var found bool
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	wg.Add(n)
	go F016Search(&wg, slice, 0, 5, target, ctx, &found)
	go F016Search(&wg, slice, 6, len(slice)-1, target, ctx, &found)
	wg.Wait()

	if found {
		fmt.Println("Found it!")
	} else if ctx.Err() != nil {
		fmt.Println("Timeout！Not Found")
	} else {
		fmt.Println("Not Found")
	}
}
func F016Search(wg *sync.WaitGroup, slice []int, i, j, target int, ctx context.Context, found *bool) {
	defer wg.Done()
	for x := i; x <= j; x++ {
		select {
		case <-ctx.Done():
			return
		default:
		}
		if slice[x] == target {
			*found = true
			return
		}
	}
}
