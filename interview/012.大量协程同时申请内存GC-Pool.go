package interview

import (
	"bytes"
	"fmt"
	"runtime"
	"sync"
	"time"
)

var pool = sync.Pool{New: func() interface{} {
	//fmt.Println("new buffer ..")
	return new(bytes.Buffer)
}}

func F012() {
	runtime.GOMAXPROCS(1) // 没用，设了1也会OOM
	go func() {
		for {
			processRequest(1 << 28) // 256MiB
		}
	}()
	for i := 0; i < 20; i++ {
		go func() { // 这个地方注释掉,就不会OOM了
			for {
				processRequest(1 << 10) // 1KiB
			}
		}()
	}
	var stats runtime.MemStats
	for i := 0; ; i++ {
		runtime.ReadMemStats(&stats)
		fmt.Printf("Cycle %d: %dB\n", i, stats.Alloc)
		time.Sleep(time.Second)
		//runtime.GC()
	}
}

func processRequest(size int) {
	b := pool.Get().(*bytes.Buffer)
	time.Sleep(200 * time.Millisecond)
	b.Grow(size)
	pool.Put(b)
	time.Sleep(10 * time.Millisecond)
}
