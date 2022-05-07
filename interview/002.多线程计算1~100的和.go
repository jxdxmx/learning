package interview

import (
	"fmt"
	"sync"
)

//多线程计算1~100的和
//思路：
//	waitGroup协调多线程
//	1~100分段计算，然后一个线程算总和
//	生产者消费者模式

func F0002() {
	var elem = make(chan int)
	var result = make(chan int)
	var wg sync.WaitGroup
	var sum = func(s, e int) {
		defer wg.Done()
		ans := 0
		for i := s; i <= e; i++ {
			ans += i
		}
		elem <- ans
	}
	wg.Add(4)
	go sum(1, 25)
	go sum(26, 50)
	go sum(51, 75)
	go sum(76, 100)
	go func() {
		var ans = 0
		for e := range elem {
			ans += e
		}
		result <- ans
	}()
	wg.Wait()
	close(elem)
	fmt.Println(<-result)
}
