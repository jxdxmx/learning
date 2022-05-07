package interview

import (
	"fmt"
	"sync"
)

//交替打印数字和字母
//问题描述
//使用两个goroutine 交替打印序列,一个 goroutine 打印数字,另外一个 goroutine 打印字母,最终效果如下:
//12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728

//解题思路
//问题很简单，使⽤ channel 来控制打印的进度。使⽤两个 channel ，来分别控制数字和
//字⺟的打印序列， 数字打印完成后通过 channel 通知字⺟打印, 字⺟打印完成后通知数
//字打印，然后周⽽复始的⼯作。
//源码参考
//
//这⾥⽤到了两个 channel 负责通知，letter负责通知打印字⺟的goroutine来打印字⺟，
//number⽤来通知打印数字的goroutine打印数字。
//wait⽤来等待字⺟打印完成后退出循环。
func F0001() {
	var ch123 = make(chan struct{}, 1)
	var chAbc = make(chan struct{}, 1)
	var chEnd = make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		i := 0
		for {
			select {
			case <-ch123:
				fmt.Print(i + 1)
				fmt.Print(i + 2)
				i += 2
				chAbc <- struct{}{}
			}
			select {
			case <-chEnd:
				return
			default:
			}
		}
	}()
	go func() {
		defer wg.Done()
		j := 0
		var rs = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
		for {
			select {
			case <-chAbc:
				fmt.Print(string(rs[j]))
				fmt.Print(string(rs[j+1]))
				j += 2
				ch123 <- struct{}{}
				if j >= len(rs)-1 {
					close(chEnd)
					return
				}
			}
		}
	}()
	ch123 <- struct{}{}
	wg.Wait()
}
