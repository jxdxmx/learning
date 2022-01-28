package main

import (
	"fmt"
)

// f20211217

//type sp interface {
//	Out(key string, val interface{})                  //存入key /val，如果该key读取的 goroutine挂起，则唤醒。此方法不会阻塞，时刻都可以立即执行并返回
//	Rd(key string, timeout time.Duration) interface{} //读取一个key，如果 key不存在阻塞，等待key存在或者超时
//}
//
//type SMap struct {
//	m   map[string]interface{}
//	chs map[string]chan struct{}
//	mmu sync.Mutex
//	cmu sync.Mutex
//}

//func f20211217() {
//	sm := newMap()
//	for i := 0; i < 10000; i++ {
//		key1 := "name-" + strconv.Itoa(i)
//		go func(x string) {
//			sm.Out(x, "xuzg"+strconv.Itoa(i))
//		}(key1)
//		go func(x string) {
//			log.Println(sm.Rd(x, time.Second*1))
//		}(key1)
//		time.Sleep(time.Nanosecond)
//	}
//	//go func() {
//	//	time.Sleep(time.Second * 2)
//	//	//sm.Out("name2", "name2")
//	//}()
//	//log.Println(sm.Rd("name2", time.Second*3))
//}
//func newMap() SMap {
//	sm := SMap{}
//	sm.m = make(map[string]interface{})
//	sm.chs = make(map[string]chan struct{})
//	sm.mmu = sync.Mutex{}
//	sm.cmu = sync.Mutex{}
//	return sm
//}
//func (s *SMap) Out(key string, val interface{}) {
//	s.mmu.Lock()
//	s.m[key] = val
//	s.mmu.Unlock()
//	s.cmu.Lock()
//	ch := s.chs[key]
//	s.cmu.Unlock()
//	if ch != nil {
//		close(ch)
//	}
//}
//func (s *SMap) Rd(key string, timeout time.Duration) (interface{}, error) {
//	s.mmu.Lock()
//	val, ok := s.m[key]
//	s.mmu.Unlock()
//	if !ok {
//		s.cmu.Lock()
//		ch := s.chs[key]
//		if ch == nil {
//			ch = make(chan struct{})
//			s.chs[key] = ch
//		}
//		s.cmu.Unlock()
//		select {
//		case <-ch:
//			s.mmu.Lock()
//			val, _ = s.m[key]
//			s.mmu.Unlock()
//			return val, nil
//		case <-time.After(timeout):
//			return nil, errors.New("time out")
//		}
//	}
//	return val, nil
//}

//func f20211217() {
//	var wg sync.WaitGroup
//	var ch = make(chan int64)
//	wg.Add(2)
//	go func() {
//		defer wg.Done()
//		for i := 0; i < 5; i++ {
//			max := big.NewInt(100)
//			r, _ := rand.Int(rand.Reader, max)
//			ch <- r.Int64()
//		}
//		close(ch)
//	}()
//	go func() {
//		defer wg.Done()
//		for n := range ch {
//			fmt.Println(n)
//		}
//	}()
//	wg.Wait()
//}

//type People interface {
//	Show()
//}
//type Student struct{}
//
//func (stu *Student) Show() {
//}
//func live() People {
//	var stu *Student
//	return stu
//}
//func f20211217() {
//	var x = live()
//	if x == nil {
//		fmt.Println("AAAAAAA")
//	} else {
//		fmt.Println("BBBBBBB")
//	}
//	var stu *Student
//	if stu == nil {
//		fmt.Println("AAAAAAA")
//	} else {
//		fmt.Println("BBBBBBB")
//	}
//	var inf interface{} = stu
//	if inf == nil {
//		fmt.Println("AAAAAAA")
//	} else {
//		fmt.Println("BBBBBBB")
//	}
//}

//type People interface {
//	Speak(string) string
//}
//type Student struct{}
//
//func (stu *Student) Speak(think string) (talk string) {
//	if think == "bitch" {
//		talk = "You are a good boy"
//	} else {
//		talk = "hi"
//	}
//	return
//}
//func f20211217() {
//	var s = Student{}
//	fmt.Println(s.Speak("a"))
//
//	//var peo People = Student{}
//	//think := "bitch"
//	//fmt.Println(peo.Speak(think))
//}

func f202112171111() {
	//defer_call()
	//fmt.Println("main 正常结束")
	zhoujielun(&student{
		Name: "xuzg",
	})
}

//type Student struct {
//	name string
//}

//func f202112171() {
//	m := map[string]Student{"people": {"zhoujielun"}}
//	m["people"].name = "wuyanzu"
//}
//func f20211217() {
//	runtime.GOMAXPROCS(2)
//	wg := sync.WaitGroup{}
//	wg.Add(10)
//	for i := 0; i < 5; i++ {
//		go func() {
//			fmt.Println("i: ", i)
//			wg.Done()
//		}()
//	}
//	for i := 0; i < 5; i++ {
//		go func(i int) {
//			fmt.Println("i: ", i)
//			wg.Done()
//		}(i)
//	}
//	wg.Wait()
//
//	//var i byte = 255
//	//fmt.Println(i + 1)
//	//go func() {
//	//	for i = 0; i <= 255; i++ {
//	//		fmt.Println(i)
//	//	}
//	//}()
//	//fmt.Println("Dropping mic")
//	//// Yield execution to force executing other goroutines
//	//runtime.Gosched()
//	//runtime.GC()
//	//fmt.Println("Done")
//}

//
//func defer_call() {
//	defer func() {
//		fmt.Println("defer: panic 之前1, 捕获异常")
//		if err := recover(); err != nil {
//			fmt.Println(err)
//		}
//	}()
//	defer func() { fmt.Println("defer: panic 之前2, 不捕获") }()
//
//	panic("异常内容") //触发defer出栈
//
//	defer func() { fmt.Println("defer: panic 之后, 永远执行不到") }()
//}

type student struct {
	Name string
}

func zhoujielun(v interface{}) {
	switch msg := v.(type) {
	case *student, student:
		fmt.Println(msg)
	}
}
