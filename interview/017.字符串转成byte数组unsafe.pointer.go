package interview

import (
	"fmt"
	"reflect"
	"unsafe"
)

//字符串转成byte数组,会发生内存拷贝吗
//字符串转成切片,会产生拷贝.严格来说,只要是发生类型强转都会发生内存拷贝.

func F017() {
	var s = "hello world"
	ssh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	//type StringHeader struct {
	//	Data uintptr
	//	Len  int
	//}

	bs := (*[]byte)(unsafe.Pointer(ssh))
	fmt.Println(len(*bs), string(*bs)) // 虽然转成byte切片成功，但是cap字段并没有被赋值
	//type SliceHeader struct {
	//	Data uintptr
	//	Len  int
	//	Cap  int
	//}

	a := "aaa"
	ssh2 := *(*reflect.StringHeader)(unsafe.Pointer(&a))
	b := *(*[]byte)(unsafe.Pointer(&ssh2))
	fmt.Printf("%v", b)
}
