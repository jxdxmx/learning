package interview

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// F009 为sync.WaitGroup⽀持WaitTimeout功能
func F009() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(time.Second * 5)
		wg.Done()
	}()
	fmt.Println("start wait")
	F009WgWaitTimeout(&wg, 30*time.Second)
	fmt.Println("end wait")
}

func F009WgWaitTimeout(wg *sync.WaitGroup, timeout time.Duration) {
	ctx, _ := context.WithTimeout(context.Background(), timeout)
	var done = make(chan struct{}, 1)
	go func() {
		wg.Wait()
		done <- struct{}{}
	}()
	select {
	case <-ctx.Done():
	case <-done:
	}
}
