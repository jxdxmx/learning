package interview

import "fmt"

//写出以下打印结果，并解释下为什么这么打印的。
// a b new
// a b new

func F011() {
	str1 := []string{"a", "b", "c"}
	str2 := str1[1:]
	str2[1] = "new"
	fmt.Println(str1)
	str2 = append(str2, "z", "x", "y")
	fmt.Println(str1)
}
