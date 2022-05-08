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

	//原因：size为1<<28的goroutine大概率会Get到1<<10的bytes.Buffer，
	//然后再grow，也就是把1k的那个缓存grow到了256M。那自然了，慢慢慢慢的，
	//大量10kbuffer被grow成了256M，内存就暴涨了。
	//由于goroutine太多，所以没法GC，仍然会残留大量对象在pool中。因此会OOM。
	//如果把【processRequest(1 << 10) // 1KiB】去掉，就不会暴涨。这是因为，每次都能获取到缓存的对象，不存在从1k>>256M的情形。
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
