package interview

import "fmt"

func F013() {
	var ch chan int
	var count int
	go func() {
		ch <- count
	}()
	go func() {
		count++
		close(ch) // panic: close of nil channel
	}()
	<-ch
	fmt.Println(count)
}
