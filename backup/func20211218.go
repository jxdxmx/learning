package backup

import (
	"fmt"
	"sync"
	"time"
)

// f20211218
type data struct {
	sync.Mutex
}

func (d *data) test(s string) {
	d.Lock()
	defer d.Unlock()
	for i := 0; i < 5; i++ {
		fmt.Println(s, i)
		time.Sleep(time.Second)
	}
}

func f20211218() {
	var wg sync.WaitGroup
	wg.Add(2)
	var d data
	go func() {
		defer wg.Done()
		d.test("read")
	}()
	go func() {
		defer wg.Done()
		d.test("write")
	}()
	wg.Wait()
}

//
//type data struct {
//	*sync.Mutex
//}
//
//func (d data) test(s string) {
//	d.Lock()
//	defer d.Unlock()
//
//	for i := 0; i < 5; i++ {
//		fmt.Println(s, i)
//		time.Sleep(time.Second)
//	}
//}
//
//func f20211218() {
//	var wg sync.WaitGroup
//	wg.Add(2)
//	var d = data{Mutex: &sync.Mutex{}}
//
//	go func() {
//		defer wg.Done()
//		d.test("read")
//	}()
//
//	go func() {
//		defer wg.Done()
//		d.test("write")
//	}()
//
//	wg.Wait()
//}

//type T struct{}
//
//func (*T) foo() {
//}
//func (T) bar() {
//}
//
//type S struct {
//	*T
//}
//
//func f20211218() {
//	s := S{}
//	_ = s.foo
//	s.foo()
//	_ = s.bar
//}

//type T struct {
//	x int
//	y *int
//}

//func f20211218() {
//	var p *[5]string
//	for i, v := range p {
//		_ = len(p)
//		fmt.Println(i, v)
//	}
//
//	//x := make([]int, 2, 10)
//	//_ = x[6:10]
//	//_ = x[6:7]
//	//_ = x[2:]
//	//
//	//i := 20
//	//t := T{10, &i}
//	//p := &t.x
//	//*p++
//	//*p--
//	//t.y = p
//	//fmt.Println(*t.y)
//}

//func f20211218() {
//	var x interface{}
//	var y interface{} = []int{3, 5}
//	_ = x == x
//	_ = x == y
//	_ = y == y
//}
//= make(map[string]int)
//type N int
//
//func (n *N) test() {
//	fmt.Println(*n)
//}
//func f20211218() {
//	var n N = 10
//	p := &n
//	n++
//	f1 := n.test
//	n++
//	f2 := p.test
//	n++
//	fmt.Println(n)
//	f1()
//	f2()
//}

//func f20211218() {
//	s := make([]int, 3, 9)
//	fmt.Println(len(s))
//	s2 := s[4:90]
//	fmt.Println(len(s2))
//}

//type T struct {
//	n int
//}
//
//func getT() T {
//	return T{}
//}
//
//func f20211218() {
//	getT().n = 1
//}

//type N int
//
//func (n N) test(add int) {
//	fmt.Println(n + N(add))
//}
//func f20211218() {
//	var n N = 10
//	fmt.Println(n)
//
//	n++
//	f1 := N.test
//	f1(n, 100)
//
//	n++
//	f2 := (*N).test
//	f2(&n, 1000)
//}

//type N int
//
//func (n N) value() {
//	n++
//	fmt.Printf("v:%p,%v\n", &n, n)
//}
//
//func (n *N) pointer() {
//	*n++
//	fmt.Printf("v:%p,%v\n", n, *n)
//}
//
//func f20211218() {
//	var a N = 25
//	p := &a
//	p1 := &p
//	p1.value()
//	p1.pointer()
//}

//type T struct {
//	n int
//}
//
//func (t *T) Set(n int) {
//	t.n = n
//}
//func getT() T {
//	return T{}
//}
//func f20211218() {
//	t := getT()
//	t.Set(1)
//}

//func f20211218() {
//	var fn1 = func() {}
//	var fn2 = func() {}
//	if fn1 != fn2 {
//		println("fn1 not equal fn2")
//	}
////}
//type X struct {
//	age int
//}
//
//func (x *X) test() {
//	println(x)
//}
//func f20211218() {
//	var a *X
//	a.test()
//	fmt.Println(a.age)
//	//var b = X{}
//	//b.test()
//	//X{}.test()
//}

//type T struct {
//	ls []int
//}
//
//func foo(t T) {
//	t.ls[0] = 100
//	t.ls = append(t.ls, 5, 6)
//}
//
//func f20211218() {
//	var t = T{
//		ls: []int{1, 2, 3},
//	}
//
//	foo(t) // 切片是引用类型
//	fmt.Println(t.ls[0])
//}

//func f20211218() {
//	//const x = 123
//	//const y = 1.23 // 不报错
//	//var z = 1.23   // 报错！
//	//fmt.Println(x)
//	fmt.Println(^2) // 0000010 11111101
//
//}

//func test(x byte) {
//	fmt.Println(x)
//}
//
//func f20211218() {
//	var a byte = 0x11
//	var b uint8 = a
//	var c uint8 = a + b
//	test(c)
//}

//func f20211218() {
//	x := interface{}(nil)
//	_ = x
//	y := (*int)(nil)
//	a := y == x
//	b := y == nil
//	d, c := x.(interface{})
//	println(a, b, c, d) // false true false
//	var e interface{} = 1
//	f, g := e.(interface{})
//	println(f, g)
//
//	var s []int
//	s = append(s, 1)
//	var m = make(map[string]int)
//	m["one"] = 1
//
//}

//type ConfigOne struct {
//	Daemon string
//}
//
////func (c *ConfigOne) String() string {
////	//return fmt.Sprintf("print: %v", c)
////}
//func f20211218() {
//	c := &ConfigOne{}
//	//c.String()
//	fmt.Printf("%+v", c)
//}

//
//var x = []int{2: 2, 3, 0: 1}
//
//func f20211218() {
//	fmt.Println(x)
//}

//type Param map[string]interface{}
//type Show struct {
//	*Param
//}
//
//func main() {
//	s := new(Show)
//	s.Param["day"] = 2
//	//(*s.Param)["day"] = 2
//}

//func Stop(stop <-chan bool) {
//	close(stop)
//}

//func Foo(x interface{}) {
//	if x == nil {
//		fmt.Println("empty interface")
//		return
//	}
//	fmt.Println("non-empty interface")
//}
//func f20211218() {
//	var x *int
//	Foo(x)
//}

//type User struct{}
//type User1 User
//type User2 = User
//
//func (i User1) m1() {
//	fmt.Println("m1")
//}
//func (i User) m2() {
//	fmt.Println("m2")
//}
//
//func f20211218() {
//	var i1 User1
//	var i2 User2
//	i1.m1()
//	i2.m2()
//}

//const i = 100
//
//var j = 123
//
//func f20211218() {
//	fmt.Println(&i, i)
//	fmt.Println(&j, j)
//}

//func f20211218() {
//	var arr = make([]int)
//	var a = [...]int{1, 2, 3, 4, 5}
//	var r [5]int
//
//	for i, v := range a {
//		if i == 0 {
//			a[1] = 12
//			a[2] = 13
//		}
//		r[i] = v
//	}
//	fmt.Println("r = ", r)
//	fmt.Println("a = ", a)
//}

//func f20211218() {
//	var a = []int{1, 2, 3, 4, 5}
//	var r [5]int
//
//	for i, v := range a {
//		if i == 0 {
//			a[1] = 12
//			a[2] = 13
//		}
//		r[i] = v
//	}
//	fmt.Println("r = ", r)
//	fmt.Println("a = ", a)
//}

//func f(n int) (r int) {
//	defer func() {
//		r += n
//		recover()
//	}()
//	var f func()
//	defer f()
//	f = func() {
//		r += 2
//	}
//	return n + 1
//}
//
//func f20211218() {
//	fmt.Println(f(3))
//}

//type Direction int
//
//const (
//	North Direction = iota
//	East
//	South
//	West
//)
//
//func (d Direction) String() string {
//	return [...]string{"North", "East", "South", "West"}[d]
//}
//
//func f20211218() {
//	//fmt.Println(South)
//	v := []int{1, 2, 3}
//	for i := range v {
//		v = append(v, i)
//	}
//	fmt.Println(v)
//}

//type S struct {
//}
//
//func f(x interface{}) {
//}
//
//func g(x *interface{}) {
//}
//
//func f20211218() {
//	s := S{}
//	p := &s
//	f(s) //A
//	g(s) //B
//	f(p) //C
//	g(p) //D
//}

//
//func increaseA() int {
//	var i int
//	defer func() {
//		i++
//	}()
//	return i
//}
//
//func increaseB() (r int) {
//	defer func() {
//		r++
//	}()
//	return r
//}
//
//func f20211218() {
//	fmt.Println(increaseA())
//	fmt.Println(increaseB())
//}

//func f20211218() {
//	var a = []string{"1", "a"}
//	var l = len(a)
//	var arr = (*[l]string)(a)
//	arr[0] = "b"
//	fmt.Println(a, arr) // [b a] &[b a]
//}

//func f20211218() {
//	s := [3]int{1, 2, 3}
//	a := s[:0]
//	b := s[:2]
//	c := s[1:2:cap(s)]
//	d := s[:]
//	e := s[1:1:3]
//	fmt.Println(cap(s), cap(a), cap(b), cap(c), cap(d), cap(e))
//}

//func f20211218() {
//	var i = 97
//	var a = 99988333
//	var r rune = 48
//	fmt.Println(string(i), string(a), string(r))
//	fmt.Println(int("2"))
//	strconv.Atoi()
//}

//func f20211218() {
//	var s1 []int
//	//var s2 = []int{}
//	if s1 == nil {
//		fmt.Println("yes nil")
//	} else {
//		fmt.Println("no nil")
//	}
//}

//func hello(i int) {
//	fmt.Println(i)
//}
//func f20211218() {
//	i := 5
//	defer hello(i) // 5
//	i = i + 10
//	str := "hello"
//	str[0] = 'x'
//	fmt.Println(str)
//
//}

//func f20211218() {
//	i := -5
//	j := +5
//	fmt.Printf("%+d %+d", i, j)
//}

//func f20211218() {
//	var m = make(map[int]int)
//	a := [2]int{1, 2}
//	b := [3]int{1, 2}
//	fmt.Println(a, b, cap(a), cap(m))
//}

//
//type person struct {
//	name string
//}
//func f20211218() {
//	//var m map[person]int
//	//p := person{"mike"}
//	//m[p] = 3
//	//fmt.Println(m[p])
//	a := [2]int{1, 2}
//	b := [3]int{1, 2}
//	//fmt.Println(a, b, a == b)
//	fmt.Println(a, b, cap(a))
//}

//func hello() []string {
//	return nil
//}

//func f20211218() {
//	var m map[string]int
//	m["23"] = 3
//	fmt.Println(m["23"])
//	//h := hello()
//	//if h == nil {
//	//	fmt.Println("nil")
//	//} else {
//	//	fmt.Println("not nil")
//	//}
//}

//type MyInt1 int
//type MyInt2 = int
//
//const (
//	x = iota
//	_
//	y
//	z = "zz"
//	k
//	p = iota
//	q = 'a'
//	o
//)
//func init() {
//
//}
//func f20211218() {
//	//fmt.Println(x, y, z, k, p, q, o)
//	init()
//}

//func f20211218() {
//	//var i int = 0
//	//var i1 MyInt1 = i
//	//var i2 MyInt2 = i
//	//fmt.Println(i1, i2)
//	var b byte = 5
//	var r rune = rune(b)
//	var ui8 uint8 = 5
//	b = ui8
//}

//func f20211218() {
//	sn1 := struct {
//		age  int
//		name string
//	}{age: 11, name: "qq"}
//	sn2 := struct {
//		age  int
//		name string
//	}{age: 11, name: "qq"}
//	if sn1 == sn2 {
//		fmt.Println("sn1 == sn2")
//	}
//	//sm1 := struct {
//	//	age int
//	//	m   map[string]string
//	//}{age: 11, m: map[string]string{"a": "1"}}
//	//sm2 := struct {
//	//	age int
//	//	m   map[string]string
//	//}{age: 11, m: map[string]string{"a": "1"}}
//	//if sm1 == sm2 {
//	//	fmt.Println("sm1 == sm2")
//	//}
//}

//func f20211218() {
//	//var m sync.Map
//	//m.LoadOrStore("a", 1)
//	//m.Delete("a")
//	//fmt.Println(len(m))
//	//m.ra
//	fmt.Println([...]string{"1"} == [...]string{"1"})
//}

//func f20211218() {
//	var ch chan int
//	var ch1 = ch
//	ch = make(chan int)
//	fmt.Println("ch : ", ch)
//	fmt.Println("ch1 : ", ch1)
//var ch chan int
//go func() {
//	defer fmt.Println("go1 end")
//	ch = make(chan int, 1)
//	fmt.Println("go1:", ch)
//	ch <- 1
//	fmt.Println("go 1")
//}()
//go func(ch chan int) {
//	defer fmt.Println("go2 end")
//	time.Sleep(time.Second * 2)
//	fmt.Println("go2:", ch)
//	fmt.Println("go 2-1")
//	fmt.Println(<-ch)
//	fmt.Println("go 2-2")
//}(ch)
//c := time.Tick(1 * time.Second)
//for range c {
//	fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
//}
//}

//var pool = sync.Pool{New: func() interface{} {
//	return new(bytes.Buffer)
//}}
//
//func f20211218() {
//	go func() {
//		for {
//			processRequest(1 << 28) // 256MiB
//		}
//	}()
//	for i := 0; i < 1000; i++ {
//		go func() {
//			for {
//				processRequest(1 << 10) // 1KiB
//			}
//		}()
//	}
//	var stats runtime.MemStats
//	for i := 0; ; i++ {
//		runtime.ReadMemStats(&stats)
//		fmt.Printf("Cycle %d: %dB\n", i, stats.Alloc)
//		time.Sleep(time.Second)
//		runtime.GC()
//	}
//}
//func processRequest(size int) {
//	b := pool.Get().(*bytes.Buffer)
//	time.Sleep(200 * time.Millisecond)
//	b.Grow(size)
//	pool.Put(b)
//	time.Sleep(1 * time.Millisecond)
//}

//type MyMutex struct {
//	count int
//	sync.Mutex
//}
//
//func f20211218() {
//	var mu MyMutex
//	mu.Lock()
//	var mu2 = mu
//	mu.count++
//	mu.Unlock()
//	mu2.Lock()
//	mu2.count++
//	mu2.Unlock()
//	fmt.Println(mu.count, mu2.count)
//}

//func f20211218() {
//	var wg sync.WaitGroup
//	wg.Add(1)
//	go func() {
//		//time.Sleep(time.Millisecond) // 这里加这条语句就是为了保证wait方法在下面的wg.Add前执行到
//		wg.Done()
//		//time.Sleep(time.Nanosecond) // 如果加上这条语句，那么不会报错！将正常退出
//		wg.Add(1)
//	}()
//	time.Sleep(time.Second)
//	wg.Wait()
//}

//var mu sync.RWMutex
//var count int
//
//func f20211218() {
//	go A()
//	time.Sleep(2 * time.Second)
//	mu.Lock()
//	defer mu.Unlock()
//	count++
//	fmt.Println(count)
//}
//func A() {
//	mu.RLock()
//	defer mu.RUnlock()
//	B()
//}
//func B() {
//	time.Sleep(1 * time.Second)
//	C()
//}
//func C() {
//	mu.RLock()
//	defer mu.RUnlock()
//}

//func main() {
//	var a = [1]string{"x"}
//	a = append(a, "b")
//	var x = [...]string{"a", "b", "c"}
//	x = append(x, "d")
//	var y = []string{"a"}
//	var z = []string{"a"}
//	fmt.Println([...]string{"1"} == [...]string{"1"})
//	fmt.Println([]string{"1"} == []string{"1"})
//	fmt.Println(y == z)
//}

//const (
//	a = iota
//	b = iota
//)
//const (
//	name  = "menglu"
//	name2 = "menglu2"
//	c     = iota
//	d     = iota
//)
//
//func f20211218() {
//	fmt.Println(a)
//	fmt.Println(b)
//	fmt.Println(c)
//	fmt.Println(d)
//}

//func f20211218() {
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
//		select {
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
