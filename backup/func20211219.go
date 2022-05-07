package backup

import "fmt"

// f20211219
func f20211219() {
	var x interface{}
	var y interface{} = []int{3, 5}
	fmt.Println(x == x)
	fmt.Println(x == y)
	fmt.Println(y == y)
}

//func f20211219() {
//	var bs = make([]byte, 100000) // 注意,这里必须初始化好bs的长度，否则最终得到的是空字节切片
//	runtime.Stack(bs, true)
//	fmt.Println("before task1 ...")
//	fmt.Println(string(bs))
//	task1()
//	time.Sleep(2 * time.Second)
//	runtime.Stack(bs, true)
//	fmt.Println("after task1 ...")
//	fmt.Println(string(bs))
//}
//
//func task1() {
//	var wg sync.WaitGroup
//	wg.Add(1)
//	go func() {
//		defer wg.Done()
//		time.Sleep(time.Second * 2)
//		fmt.Println("wg end")
//	}()
//	var finish = make(chan struct{}, 1)
//	go func() {
//		wg.Wait()
//		select { // 这里使用select来防止泄露
//		case finish <- struct{}{}:
//		default:
//		}
//	}()
//	go func() {
//		time.AfterFunc(time.Second*1, func() {
//			fmt.Println("timeout ...")
//			finish <- struct{}{}
//		})
//	}()
//	<-finish
//}
