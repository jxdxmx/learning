package interview

import (
	"fmt"
	"time"
)

func F008() {
	ch := time.Tick(time.Second)
	i := 1
	for {
		select {
		case <-ch:
			fmt.Print(i, ",")
			if i%20 == 0 {
				fmt.Println()
			}
			i++
		}
	}
}
