package backup

import (
	"fmt"
	"sync"
)

func f20220302() {
	var result = 0
	var sChan = make(chan int)
	var wg = &sync.WaitGroup{}
	wg.Add(4)
	go sum(1, 25, sChan, wg)
	go sum(26, 50, sChan, wg)
	go sum(51, 75, sChan, wg)
	go sum(76, 100, sChan, wg)
	go func() {
		for s := range sChan {
			result += s
		}
	}()
	wg.Wait()
	close(sChan)
	fmt.Println(result)
}

func sum(s, e int, sChan chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	rst := 0
	for i := s; i <= e; i++ {
		rst += i
	}
	sChan <- rst
}
