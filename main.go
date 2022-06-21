package main

import "fmt"

type Per struct {
	A int
}

func main() {
	var i = 0
	for ; i < 5; i++ {
		if i < 5 {
			//break
		}
	}
	fmt.Println(i)
}
