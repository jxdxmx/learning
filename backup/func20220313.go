package backup

import (
	crand "crypto/rand"
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

func randInt64() int64 {
	bi := big.NewInt(time.Now().UnixNano())
	cr, _ := crand.Int(crand.Reader, bi)
	j := cr.Int64()
	rand.Seed(j)
	return rand.Int63()
}

func f20220313() {
	//var slice1 = make([]int, 2)
	//var slice2 = make([]int, 2)
	//fmt.Println(slice1 == slice2)

	//var m =make(map[chan int]int)
	var ch1 = make(chan int, 2)
	var ch2 = make(chan int, 2)
	fmt.Println(ch1 == ch2)

	//var s = student{
	//	Name: "xuzhigang",
	//}
	var m2 = make(map[student]int)
	fmt.Println(m2)

	var int1, int0 = 0, 0
	//for i := 0; i < 10000; i++ {
	//	j := randInt64()
	//	if j&1 == 1 {
	//		int1++
	//	} else {
	//		int0++
	//	}
	//}
	//fmt.Println(int1, int0)
	//
	//int1, int0 = 0, 0
	for i := 0; i < 10000; i++ {
		rand.Seed(time.Now().UnixNano() + int64(i))
		j := rand.Int()
		if j&1 == 1 {
			int1++
		} else {
			int0++
		}
		//fmt.Print(rand.Intn(10), " ")
	}
	fmt.Println(int1, int0)
	//fmt.Println()
	//rand.Seed(50000)
	//for i := 0; i < 10; i++ {
	//	fmt.Print(rand.Intn(10), " ")
	//}

	int1, int0 = 0, 0
	for i := 0; i < 10000; i++ {
		bi := big.NewInt(50000)
		cr, _ := crand.Int(crand.Reader, bi)
		j := cr.Int64()
		if j&1 == 1 {
			int1++
		} else {
			int0++
		}
	}
	fmt.Println(int1, int0)
}
