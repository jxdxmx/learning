package main

import "fmt"

type Per struct {
	A int
}

func main() {
	fmt.Println(len("aaa"))
	str := "abab"
	fmt.Println(str[0:2] == str[2:4])
}
